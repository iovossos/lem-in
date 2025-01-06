import os
os.environ['PYGAME_HIDE_SUPPORT_PROMPT'] = "hide"
import sys
import pygame
import math

WINDOW_WIDTH = 1200
WINDOW_HEIGHT = 800
ROOM_RADIUS = 20
FPS = 60

# Colors
WHITE = (255, 255, 255)
BLACK = (0, 0, 0)
RED = (255, 0, 0)
GREEN = (0, 255, 0)
BLUE = (0, 0, 255)

class Room:
    def __init__(self, name, x, y, is_start=False, is_end=False):
        self.name = name
        self.x = int(x)
        self.y = int(y)
        self.scaled_x = 0
        self.scaled_y = 0
        self.is_start = is_start
        self.is_end = is_end

class Ant:
    def __init__(self, id, room):
        self.id = id
        self.current_room = room
        self.target_room = room
        self.x = room.scaled_x
        self.y = room.scaled_y
        self.progress = 1.0  # Animation progress
        self.history = []  # Track room history

class Visualizer:
    def __init__(self):
        pygame.init()
        self.screen = pygame.display.set_mode((WINDOW_WIDTH, WINDOW_HEIGHT))
        pygame.display.set_caption("Lem-in Visualizer")
        self.font = pygame.font.Font(None, 24)
        self.clock = pygame.time.Clock()
        self.rooms = {}
        self.connections = []
        self.moves = []
        self.ants = {}
        self.current_move = 0
        self.paused = True
        self.animation_speed = 0.02  # Controls animation speed
        self.time_since_last_move = 0
        self.move_delay = 1000  # Milliseconds between moves
        self.start_room = None  # Track start room

    def show_error(self, error_message):
        self.screen.fill(WHITE)
        # Split error message into lines
        lines = error_message.strip().split('\n')
        y = WINDOW_HEIGHT // 2 - (len(lines) * 30) // 2
        
        for line in lines:
            text = self.font.render(line, True, RED)
            text_rect = text.get_rect(center=(WINDOW_WIDTH // 2, y))
            self.screen.blit(text, text_rect)
            y += 30
            
        pygame.display.flip()
        # Wait for a moment to show error
        pygame.time.wait(3000)

    def parse_input(self):
        try:
            # Read all input
            input_text = sys.stdin.read()
            
            # Check for error messages
            if 'error' in input_text.lower() or 'exit status' in input_text.lower():
                self.show_error(input_text)
                return False
                
            # Split into lines and continue with normal parsing
            lines = input_text.strip().split('\n')
            self.num_ants = int(lines[0])
            i = 1
            is_start = False
            is_end = False
            start_room = None
            parsing_rooms = True

            while i < len(lines):
                line = lines[i]
                if not line:
                    i += 1
                    continue
                    
                if line == "#rooms":
                    i += 1
                    continue
                    
                if parsing_rooms:
                    if line == "##start":
                        is_start = True
                    elif line == "##end":
                        is_end = True
                    elif " " in line:
                        parts = line.split()
                        if len(parts) == 3:
                            name, x, y = parts
                            room = Room(name, x, y, is_start, is_end)
                            self.rooms[name] = room
                            if is_start:
                                self.start_room = room
                                is_start = False
                            if is_end:
                                is_end = False
                    elif "-" in line:
                        parsing_rooms = False
                        
                if not parsing_rooms:
                    if line.startswith("L"):
                        moves = []
                        for move in line.split():
                            if "-" in move:
                                ant, room = move[1:].split("-")
                                moves.append((ant, room))
                        if moves:
                            self.moves.append(moves)
                    elif "-" in line:
                        src, dst = line.split("-")
                        if src in self.rooms and dst in self.rooms:
                            self.connections.append((src, dst))
                i += 1
                
            # Initialize ants
            for i in range(1, self.num_ants + 1):
                self.ants[str(i)] = Ant(i, self.start_room)
                self.ants[str(i)].x = self.start_room.scaled_x
                self.ants[str(i)].y = self.start_room.scaled_y
            
            self.scale_coordinates()
            return True
            
        except Exception as e:
            self.show_error(str(e))
            return False

    def scale_coordinates(self):
        if not self.rooms:
            return
            
        # Find bounds
        min_x = min(r.x for r in self.rooms.values())
        max_x = max(r.x for r in self.rooms.values())
        min_y = min(r.y for r in self.rooms.values())
        max_y = max(r.y for r in self.rooms.values())
        
        # Calculate scale
        margin = 50
        width = max_x - min_x
        height = max_y - min_y
        scale_x = (WINDOW_WIDTH - 2 * margin) / (width if width > 0 else 1)
        scale_y = (WINDOW_HEIGHT - 2 * margin) / (height if height > 0 else 1)
        scale = min(scale_x, scale_y)
        
        # Scale coordinates
        for room in self.rooms.values():
            room.scaled_x = int((room.x - min_x) * scale + margin)
            room.scaled_y = int((room.y - min_y) * scale + margin)

    def update_ants(self):
        if self.paused:
            return

        # Update animation progress
        all_finished = True
        for ant in self.ants.values():
            if ant.progress < 1.0:
                all_finished = False
                ant.progress = min(1.0, ant.progress + self.animation_speed)
                # Interpolate position
                ant.x = ant.current_room.scaled_x + (ant.target_room.scaled_x - ant.current_room.scaled_x) * ant.progress
                ant.y = ant.current_room.scaled_y + (ant.target_room.scaled_y - ant.current_room.scaled_y) * ant.progress

        # Move to next frame if ready
        if all_finished and self.current_move < len(self.moves):
            self.time_since_last_move += self.clock.get_time()
            if self.time_since_last_move >= self.move_delay:
                self.time_since_last_move = 0
                # Store current positions before updating
                move_updates = {}
                
                # Collect all moves first
                for ant_id, room_name in self.moves[self.current_move]:
                    move_updates[ant_id] = room_name
                
                # Update all ant positions simultaneously
                for ant_id, room_name in move_updates.items():
                    ant = self.ants[ant_id]
                    ant.current_room = ant.target_room
                    ant.target_room = self.rooms[room_name]
                    ant.progress = 0.0
                    ant.history.append(room_name)
                
                self.current_move += 1

    def draw(self):
        self.screen.fill(WHITE)
        
        # Draw connections
        for src, dst in self.connections:
            start = self.rooms[src]
            end = self.rooms[dst]
            pygame.draw.line(self.screen, BLACK, 
                           (start.scaled_x, start.scaled_y),
                           (end.scaled_x, end.scaled_y), 2)
        
        # Draw rooms
        for room in self.rooms.values():
            color = GREEN if room.is_start else RED if room.is_end else BLUE
            pygame.draw.circle(self.screen, color, (room.scaled_x, room.scaled_y), ROOM_RADIUS)
            text = self.font.render(room.name, True, BLACK)
            self.screen.blit(text, (room.scaled_x - text.get_width()//2, 
                                  room.scaled_y - ROOM_RADIUS - 20))
        
        # Draw ants
        for ant in self.ants.values():
            pygame.draw.circle(self.screen, RED, (int(ant.x), int(ant.y)), ROOM_RADIUS//2)
            text = self.font.render(f"L{ant.id}", True, WHITE)
            self.screen.blit(text, (int(ant.x) - text.get_width()//2, 
                                  int(ant.y) - text.get_height()//2))
        
        # Draw info
        info = f"Move: {self.current_move}/{len(self.moves)}  {'PAUSED' if self.paused else 'PLAYING'}"
        text = self.font.render(info, True, BLACK)
        self.screen.blit(text, (10, 10))
        
        pygame.display.flip()

    def reset_ants_to_start(self):
        for ant in self.ants.values():
            ant.current_room = self.start_room
            ant.target_room = self.start_room
            ant.x = self.start_room.scaled_x
            ant.y = self.start_room.scaled_y
            ant.progress = 1.0

    def reset_ant_positions(self, move_index):
        if move_index == 0:
            self.reset_ants_to_start()
            return
            
        # Reset to positions at specific move
        for ant in self.ants.values():
            ant.history = []
        
        # Replay moves up to target index
        for i in range(move_index):
            for ant_id, room_name in self.moves[i]:
                ant = self.ants[ant_id]
                ant.current_room = ant.target_room = self.rooms[room_name]
                ant.x = ant.target_room.scaled_x
                ant.y = ant.target_room.scaled_y
                ant.progress = 1.0
                ant.history.append(room_name)

    def handle_navigation(self, direction):
        if direction == 'right' and self.current_move < len(self.moves):
            self.current_move += 1
            self.reset_ant_positions(self.current_move)
        elif direction == 'left' and self.current_move > 0:
            self.current_move -= 1
            self.reset_ant_positions(self.current_move)

    def run(self):
        # Try to parse input, exit if failed
        if not self.parse_input():
            pygame.quit()
            sys.exit(1)
        
        running = True
        while running:
            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    running = False
                elif event.type == pygame.KEYDOWN:
                    if event.key == pygame.K_SPACE:
                        self.paused = not self.paused
                    elif event.key == pygame.K_RIGHT and self.paused:
                        self.handle_navigation('right')
                    elif event.key == pygame.K_LEFT and self.paused:
                        self.handle_navigation('left')
            
            self.update_ants()
            self.draw()
            self.clock.tick(FPS)
        
        pygame.quit()

if __name__ == "__main__":
    vis = Visualizer()
    vis.run()