//+build darwin
#include <Cocoa/Cocoa.h>

@interface MenuItem : NSObject
{
  @public
    NSNumber* menuId;
    NSString* title;
    NSString* tooltip;
    short disabled;
    short checked;
}
@end

@interface AppDelegate: NSObject <NSApplicationDelegate>
  - (void)add_or_update_menu_item:(MenuItem *)item;
  - (IBAction)menuHandler:(id)sender;
  @property (assign) IBOutlet NSWindow *window;
@end
