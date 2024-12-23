import pygame
import sys
import pygame_gui
 
from pygame_gui.elements import UIButton
from pygame_gui.windows import UIColourPickerDialog
 
def build_string(colour):
    return '{' + f'R: {colour[0]}, G: {colour[1]}, B: {colour[2]}, A: {255}' + '}'

pygame.init()
 
pygame.display.set_caption('Colour Picking App')
SCREEN = pygame.display.set_mode((800, 600))
 
ui_manager = pygame_gui.UIManager((800, 600))
background = pygame.Surface((800, 600))
background.fill("#3a3b3c")
current_colour = pygame.Color(0, 0, 0)
colour_picker_label = pygame_gui.elements.UITextBox(html_text=build_string(current_colour),
                                                    relative_rect=pygame.Rect((275, 550), (250, 30)),                                                  
                                                    manager=ui_manager)
colour_picker_button = UIButton(relative_rect=pygame.Rect((25, 25), (150, 30)),
                                text='Open Colour Picker',
                                manager=ui_manager)
colour_picker = None                                    

picked_colour_surface = pygame.Surface((400, 400))
picked_colour_surface.fill(current_colour)
 
clock = pygame.time.Clock()
 
while True:
    time_delta = clock.tick(60) / 1000
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            pygame.quit()
            sys.exit()
        if event.type == pygame_gui.UI_BUTTON_PRESSED:
            if event.ui_element == colour_picker_button:
                colour_picker = UIColourPickerDialog(rect=pygame.Rect((200, 200), (400, 400)),
                                                     manager=ui_manager,
                                                     window_title='Colour Picker',
                                                     initial_colour=current_colour)
        if event.type == pygame_gui.UI_COLOUR_PICKER_COLOUR_PICKED:
            current_colour = event.colour
            colour_picker_label.html_text = build_string(current_colour)
            colour_picker_label.rebuild()
            picked_colour_surface.fill(current_colour)
        
        ui_manager.process_events(event)
        
    ui_manager.update(time_delta)
 
    SCREEN.blit(background, (0, 0))
    SCREEN.blit(picked_colour_surface, (200, 100))
 
    ui_manager.draw_ui(SCREEN)
 
    pygame.display.update()