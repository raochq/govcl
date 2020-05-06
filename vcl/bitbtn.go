
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TBitBtn struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBitBtn(owner IComponent) *TBitBtn {
    b := new(TBitBtn)
    b.instance = BitBtn_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsBitBtn(obj interface{}) *TBitBtn {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TBitBtn{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsBitBtn.
func BitBtnFromInst(inst uintptr) *TBitBtn {
    return AsBitBtn(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsBitBtn.
func BitBtnFromObj(obj IObject) *TBitBtn {
    return AsBitBtn(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsBitBtn.
func BitBtnFromUnsafePointer(ptr unsafe.Pointer) *TBitBtn {
    return AsBitBtn(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (b *TBitBtn) Free() {
    if b.instance != 0 {
        BitBtn_Free(b.instance)
        b.instance, b.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBitBtn) Instance() uintptr {
    return b.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBitBtn) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBitBtn) IsValid() bool {
    return b.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (b *TBitBtn) Is() TIs {
    return TIs(b.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (b *TBitBtn) As() TAs {
//    return TAs(b.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBitBtnClass() TClass {
    return BitBtn_StaticClassType()
}

// CN: 单击。
// EN: .
func (b *TBitBtn) Click() {
    BitBtn_Click(b.instance)
}

// CN: 是否可以获得焦点。
// EN: .
func (b *TBitBtn) CanFocus() bool {
    return BitBtn_CanFocus(b.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (b *TBitBtn) ContainsControl(Control IControl) bool {
    return BitBtn_ContainsControl(b.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (b *TBitBtn) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(BitBtn_ControlAtPos(b.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (b *TBitBtn) DisableAlign() {
    BitBtn_DisableAlign(b.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (b *TBitBtn) EnableAlign() {
    BitBtn_EnableAlign(b.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (b *TBitBtn) FindChildControl(ControlName string) *TControl {
    return AsControl(BitBtn_FindChildControl(b.instance, ControlName))
}

func (b *TBitBtn) FlipChildren(AllLevels bool) {
    BitBtn_FlipChildren(b.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (b *TBitBtn) Focused() bool {
    return BitBtn_Focused(b.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (b *TBitBtn) HandleAllocated() bool {
    return BitBtn_HandleAllocated(b.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (b *TBitBtn) InsertControl(AControl IControl) {
    BitBtn_InsertControl(b.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (b *TBitBtn) Invalidate() {
    BitBtn_Invalidate(b.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (b *TBitBtn) RemoveControl(AControl IControl) {
    BitBtn_RemoveControl(b.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (b *TBitBtn) Realign() {
    BitBtn_Realign(b.instance)
}

// CN: 重绘。
// EN: Repaint.
func (b *TBitBtn) Repaint() {
    BitBtn_Repaint(b.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (b *TBitBtn) ScaleBy(M int32, D int32) {
    BitBtn_ScaleBy(b.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (b *TBitBtn) ScrollBy(DeltaX int32, DeltaY int32) {
    BitBtn_ScrollBy(b.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (b *TBitBtn) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    BitBtn_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (b *TBitBtn) SetFocus() {
    BitBtn_SetFocus(b.instance)
}

// CN: 控件更新。
// EN: Update.
func (b *TBitBtn) Update() {
    BitBtn_Update(b.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (b *TBitBtn) BringToFront() {
    BitBtn_BringToFront(b.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (b *TBitBtn) ClientToScreen(Point TPoint) TPoint {
    return BitBtn_ClientToScreen(b.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (b *TBitBtn) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return BitBtn_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (b *TBitBtn) Dragging() bool {
    return BitBtn_Dragging(b.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (b *TBitBtn) HasParent() bool {
    return BitBtn_HasParent(b.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (b *TBitBtn) Hide() {
    BitBtn_Hide(b.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (b *TBitBtn) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return BitBtn_Perform(b.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (b *TBitBtn) Refresh() {
    BitBtn_Refresh(b.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (b *TBitBtn) ScreenToClient(Point TPoint) TPoint {
    return BitBtn_ScreenToClient(b.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (b *TBitBtn) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return BitBtn_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (b *TBitBtn) SendToBack() {
    BitBtn_SendToBack(b.instance)
}

// CN: 显示控件。
// EN: Show control.
func (b *TBitBtn) Show() {
    BitBtn_Show(b.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (b *TBitBtn) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return BitBtn_GetTextBuf(b.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (b *TBitBtn) GetTextLen() int32 {
    return BitBtn_GetTextLen(b.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (b *TBitBtn) SetTextBuf(Buffer string) {
    BitBtn_SetTextBuf(b.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (b *TBitBtn) FindComponent(AName string) *TComponent {
    return AsComponent(BitBtn_FindComponent(b.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBitBtn) GetNamePath() string {
    return BitBtn_GetNamePath(b.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBitBtn) Assign(Source IObject) {
    BitBtn_Assign(b.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBitBtn) ClassType() TClass {
    return BitBtn_ClassType(b.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBitBtn) ClassName() string {
    return BitBtn_ClassName(b.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBitBtn) InstanceSize() int32 {
    return BitBtn_InstanceSize(b.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBitBtn) InheritsFrom(AClass TClass) bool {
    return BitBtn_InheritsFrom(b.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBitBtn) Equals(Obj IObject) bool {
    return BitBtn_Equals(b.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBitBtn) GetHashCode() int32 {
    return BitBtn_GetHashCode(b.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (b *TBitBtn) ToString() string {
    return BitBtn_ToString(b.instance)
}

func (b *TBitBtn) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    BitBtn_AnchorToNeighbour(b.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (b *TBitBtn) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    BitBtn_AnchorParallel(b.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (b *TBitBtn) AnchorHorizontalCenterTo(ASibling IControl) {
    BitBtn_AnchorHorizontalCenterTo(b.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (b *TBitBtn) AnchorVerticalCenterTo(ASibling IControl) {
    BitBtn_AnchorVerticalCenterTo(b.instance, CheckPtr(ASibling))
}

func (b *TBitBtn) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    BitBtn_AnchorAsAlign(b.instance, ATheAlign , ASpace)
}

func (b *TBitBtn) AnchorClient(ASpace int32) {
    BitBtn_AnchorClient(b.instance, ASpace)
}

func (b *TBitBtn) DefaultCaption() bool {
    return BitBtn_GetDefaultCaption(b.instance)
}

func (b *TBitBtn) SetDefaultCaption(value bool) {
    BitBtn_SetDefaultCaption(b.instance, value)
}

func (b *TBitBtn) GlyphShowMode() TGlyphShowMode {
    return BitBtn_GetGlyphShowMode(b.instance)
}

func (b *TBitBtn) SetGlyphShowMode(value TGlyphShowMode) {
    BitBtn_SetGlyphShowMode(b.instance, value)
}

func (b *TBitBtn) ImageWidth() int32 {
    return BitBtn_GetImageWidth(b.instance)
}

func (b *TBitBtn) SetImageWidth(value int32) {
    BitBtn_SetImageWidth(b.instance, value)
}

func (b *TBitBtn) Action() *TAction {
    return AsAction(BitBtn_GetAction(b.instance))
}

func (b *TBitBtn) SetAction(value IComponent) {
    BitBtn_SetAction(b.instance, CheckPtr(value))
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (b *TBitBtn) Align() TAlign {
    return BitBtn_GetAlign(b.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (b *TBitBtn) SetAlign(value TAlign) {
    BitBtn_SetAlign(b.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (b *TBitBtn) Anchors() TAnchors {
    return BitBtn_GetAnchors(b.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (b *TBitBtn) SetAnchors(value TAnchors) {
    BitBtn_SetAnchors(b.instance, value)
}

func (b *TBitBtn) BiDiMode() TBiDiMode {
    return BitBtn_GetBiDiMode(b.instance)
}

func (b *TBitBtn) SetBiDiMode(value TBiDiMode) {
    BitBtn_SetBiDiMode(b.instance, value)
}

func (b *TBitBtn) Cancel() bool {
    return BitBtn_GetCancel(b.instance)
}

func (b *TBitBtn) SetCancel(value bool) {
    BitBtn_SetCancel(b.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (b *TBitBtn) Caption() string {
    return BitBtn_GetCaption(b.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (b *TBitBtn) SetCaption(value string) {
    BitBtn_SetCaption(b.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (b *TBitBtn) Constraints() *TSizeConstraints {
    return AsSizeConstraints(BitBtn_GetConstraints(b.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (b *TBitBtn) SetConstraints(value *TSizeConstraints) {
    BitBtn_SetConstraints(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Default() bool {
    return BitBtn_GetDefault(b.instance)
}

func (b *TBitBtn) SetDefault(value bool) {
    BitBtn_SetDefault(b.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (b *TBitBtn) DoubleBuffered() bool {
    return BitBtn_GetDoubleBuffered(b.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (b *TBitBtn) SetDoubleBuffered(value bool) {
    BitBtn_SetDoubleBuffered(b.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBitBtn) Enabled() bool {
    return BitBtn_GetEnabled(b.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBitBtn) SetEnabled(value bool) {
    BitBtn_SetEnabled(b.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (b *TBitBtn) Font() *TFont {
    return AsFont(BitBtn_GetFont(b.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (b *TBitBtn) SetFont(value *TFont) {
    BitBtn_SetFont(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Glyph() *TBitmap {
    return AsBitmap(BitBtn_GetGlyph(b.instance))
}

func (b *TBitBtn) SetGlyph(value *TBitmap) {
    BitBtn_SetGlyph(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Layout() TButtonLayout {
    return BitBtn_GetLayout(b.instance)
}

func (b *TBitBtn) SetLayout(value TButtonLayout) {
    BitBtn_SetLayout(b.instance, value)
}

// CN: 获取模态对话框显示结果。
// EN: .
func (b *TBitBtn) ModalResult() TModalResult {
    return BitBtn_GetModalResult(b.instance)
}

// CN: 设置模态对话框显示结果。
// EN: .
func (b *TBitBtn) SetModalResult(value TModalResult) {
    BitBtn_SetModalResult(b.instance, value)
}

func (b *TBitBtn) NumGlyphs() TNumGlyphs {
    return BitBtn_GetNumGlyphs(b.instance)
}

func (b *TBitBtn) SetNumGlyphs(value TNumGlyphs) {
    BitBtn_SetNumGlyphs(b.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (b *TBitBtn) ParentDoubleBuffered() bool {
    return BitBtn_GetParentDoubleBuffered(b.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (b *TBitBtn) SetParentDoubleBuffered(value bool) {
    BitBtn_SetParentDoubleBuffered(b.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (b *TBitBtn) ParentFont() bool {
    return BitBtn_GetParentFont(b.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (b *TBitBtn) SetParentFont(value bool) {
    BitBtn_SetParentFont(b.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (b *TBitBtn) ParentShowHint() bool {
    return BitBtn_GetParentShowHint(b.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (b *TBitBtn) SetParentShowHint(value bool) {
    BitBtn_SetParentShowHint(b.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (b *TBitBtn) PopupMenu() *TPopupMenu {
    return AsPopupMenu(BitBtn_GetPopupMenu(b.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (b *TBitBtn) SetPopupMenu(value IComponent) {
    BitBtn_SetPopupMenu(b.instance, CheckPtr(value))
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (b *TBitBtn) ShowHint() bool {
    return BitBtn_GetShowHint(b.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (b *TBitBtn) SetShowHint(value bool) {
    BitBtn_SetShowHint(b.instance, value)
}

func (b *TBitBtn) Spacing() int32 {
    return BitBtn_GetSpacing(b.instance)
}

func (b *TBitBtn) SetSpacing(value int32) {
    BitBtn_SetSpacing(b.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (b *TBitBtn) TabOrder() TTabOrder {
    return BitBtn_GetTabOrder(b.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (b *TBitBtn) SetTabOrder(value TTabOrder) {
    BitBtn_SetTabOrder(b.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (b *TBitBtn) TabStop() bool {
    return BitBtn_GetTabStop(b.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (b *TBitBtn) SetTabStop(value bool) {
    BitBtn_SetTabStop(b.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBitBtn) Visible() bool {
    return BitBtn_GetVisible(b.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBitBtn) SetVisible(value bool) {
    BitBtn_SetVisible(b.instance, value)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (b *TBitBtn) SetOnClick(fn TNotifyEvent) {
    BitBtn_SetOnClick(b.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (b *TBitBtn) SetOnContextPopup(fn TContextPopupEvent) {
    BitBtn_SetOnContextPopup(b.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (b *TBitBtn) SetOnDragDrop(fn TDragDropEvent) {
    BitBtn_SetOnDragDrop(b.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (b *TBitBtn) SetOnDragOver(fn TDragOverEvent) {
    BitBtn_SetOnDragOver(b.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (b *TBitBtn) SetOnEndDrag(fn TEndDragEvent) {
    BitBtn_SetOnEndDrag(b.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (b *TBitBtn) SetOnEnter(fn TNotifyEvent) {
    BitBtn_SetOnEnter(b.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (b *TBitBtn) SetOnExit(fn TNotifyEvent) {
    BitBtn_SetOnExit(b.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (b *TBitBtn) SetOnKeyDown(fn TKeyEvent) {
    BitBtn_SetOnKeyDown(b.instance, fn)
}

func (b *TBitBtn) SetOnKeyPress(fn TKeyPressEvent) {
    BitBtn_SetOnKeyPress(b.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (b *TBitBtn) SetOnKeyUp(fn TKeyEvent) {
    BitBtn_SetOnKeyUp(b.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (b *TBitBtn) SetOnMouseDown(fn TMouseEvent) {
    BitBtn_SetOnMouseDown(b.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (b *TBitBtn) SetOnMouseEnter(fn TNotifyEvent) {
    BitBtn_SetOnMouseEnter(b.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (b *TBitBtn) SetOnMouseLeave(fn TNotifyEvent) {
    BitBtn_SetOnMouseLeave(b.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (b *TBitBtn) SetOnMouseMove(fn TMouseMoveEvent) {
    BitBtn_SetOnMouseMove(b.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (b *TBitBtn) SetOnMouseUp(fn TMouseEvent) {
    BitBtn_SetOnMouseUp(b.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (b *TBitBtn) DockClientCount() int32 {
    return BitBtn_GetDockClientCount(b.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (b *TBitBtn) DockSite() bool {
    return BitBtn_GetDockSite(b.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (b *TBitBtn) SetDockSite(value bool) {
    BitBtn_SetDockSite(b.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (b *TBitBtn) MouseInClient() bool {
    return BitBtn_GetMouseInClient(b.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (b *TBitBtn) VisibleDockClientCount() int32 {
    return BitBtn_GetVisibleDockClientCount(b.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (b *TBitBtn) Brush() *TBrush {
    return AsBrush(BitBtn_GetBrush(b.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (b *TBitBtn) ControlCount() int32 {
    return BitBtn_GetControlCount(b.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (b *TBitBtn) Handle() HWND {
    return BitBtn_GetHandle(b.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (b *TBitBtn) ParentWindow() HWND {
    return BitBtn_GetParentWindow(b.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (b *TBitBtn) SetParentWindow(value HWND) {
    BitBtn_SetParentWindow(b.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (b *TBitBtn) UseDockManager() bool {
    return BitBtn_GetUseDockManager(b.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (b *TBitBtn) SetUseDockManager(value bool) {
    BitBtn_SetUseDockManager(b.instance, value)
}

func (b *TBitBtn) BoundsRect() TRect {
    return BitBtn_GetBoundsRect(b.instance)
}

func (b *TBitBtn) SetBoundsRect(value TRect) {
    BitBtn_SetBoundsRect(b.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (b *TBitBtn) ClientHeight() int32 {
    return BitBtn_GetClientHeight(b.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (b *TBitBtn) SetClientHeight(value int32) {
    BitBtn_SetClientHeight(b.instance, value)
}

func (b *TBitBtn) ClientOrigin() TPoint {
    return BitBtn_GetClientOrigin(b.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (b *TBitBtn) ClientRect() TRect {
    return BitBtn_GetClientRect(b.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (b *TBitBtn) ClientWidth() int32 {
    return BitBtn_GetClientWidth(b.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (b *TBitBtn) SetClientWidth(value int32) {
    BitBtn_SetClientWidth(b.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (b *TBitBtn) ControlState() TControlState {
    return BitBtn_GetControlState(b.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (b *TBitBtn) SetControlState(value TControlState) {
    BitBtn_SetControlState(b.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (b *TBitBtn) ControlStyle() TControlStyle {
    return BitBtn_GetControlStyle(b.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (b *TBitBtn) SetControlStyle(value TControlStyle) {
    BitBtn_SetControlStyle(b.instance, value)
}

func (b *TBitBtn) Floating() bool {
    return BitBtn_GetFloating(b.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBitBtn) Parent() *TWinControl {
    return AsWinControl(BitBtn_GetParent(b.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBitBtn) SetParent(value IWinControl) {
    BitBtn_SetParent(b.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (b *TBitBtn) Left() int32 {
    return BitBtn_GetLeft(b.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (b *TBitBtn) SetLeft(value int32) {
    BitBtn_SetLeft(b.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (b *TBitBtn) Top() int32 {
    return BitBtn_GetTop(b.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (b *TBitBtn) SetTop(value int32) {
    BitBtn_SetTop(b.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (b *TBitBtn) Width() int32 {
    return BitBtn_GetWidth(b.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (b *TBitBtn) SetWidth(value int32) {
    BitBtn_SetWidth(b.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (b *TBitBtn) Height() int32 {
    return BitBtn_GetHeight(b.instance)
}

// CN: 设置高度。
// EN: Set height.
func (b *TBitBtn) SetHeight(value int32) {
    BitBtn_SetHeight(b.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBitBtn) Cursor() TCursor {
    return BitBtn_GetCursor(b.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBitBtn) SetCursor(value TCursor) {
    BitBtn_SetCursor(b.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (b *TBitBtn) Hint() string {
    return BitBtn_GetHint(b.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (b *TBitBtn) SetHint(value string) {
    BitBtn_SetHint(b.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBitBtn) ComponentCount() int32 {
    return BitBtn_GetComponentCount(b.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (b *TBitBtn) ComponentIndex() int32 {
    return BitBtn_GetComponentIndex(b.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (b *TBitBtn) SetComponentIndex(value int32) {
    BitBtn_SetComponentIndex(b.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBitBtn) Owner() *TComponent {
    return AsComponent(BitBtn_GetOwner(b.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBitBtn) Name() string {
    return BitBtn_GetName(b.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBitBtn) SetName(value string) {
    BitBtn_SetName(b.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBitBtn) Tag() int {
    return BitBtn_GetTag(b.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBitBtn) SetTag(value int) {
    BitBtn_SetTag(b.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (b *TBitBtn) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(BitBtn_GetAnchorSideLeft(b.instance))
}

// CN: 设置左边锚点。
// EN: .
func (b *TBitBtn) SetAnchorSideLeft(value *TAnchorSide) {
    BitBtn_SetAnchorSideLeft(b.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (b *TBitBtn) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(BitBtn_GetAnchorSideTop(b.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (b *TBitBtn) SetAnchorSideTop(value *TAnchorSide) {
    BitBtn_SetAnchorSideTop(b.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (b *TBitBtn) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(BitBtn_GetAnchorSideRight(b.instance))
}

// CN: 设置右边锚点。
// EN: .
func (b *TBitBtn) SetAnchorSideRight(value *TAnchorSide) {
    BitBtn_SetAnchorSideRight(b.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (b *TBitBtn) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(BitBtn_GetAnchorSideBottom(b.instance))
}

// CN: 设置底边锚点。
// EN: .
func (b *TBitBtn) SetAnchorSideBottom(value *TAnchorSide) {
    BitBtn_SetAnchorSideBottom(b.instance, CheckPtr(value))
}

func (b *TBitBtn) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(BitBtn_GetChildSizing(b.instance))
}

func (b *TBitBtn) SetChildSizing(value *TControlChildSizing) {
    BitBtn_SetChildSizing(b.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (b *TBitBtn) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(BitBtn_GetBorderSpacing(b.instance))
}

// CN: 设置边框间距。
// EN: .
func (b *TBitBtn) SetBorderSpacing(value *TControlBorderSpacing) {
    BitBtn_SetBorderSpacing(b.instance, CheckPtr(value))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (b *TBitBtn) DockClients(Index int32) *TControl {
    return AsControl(BitBtn_GetDockClients(b.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (b *TBitBtn) Controls(Index int32) *TControl {
    return AsControl(BitBtn_GetControls(b.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBitBtn) Components(AIndex int32) *TComponent {
    return AsComponent(BitBtn_GetComponents(b.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (b *TBitBtn) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(BitBtn_GetAnchorSide(b.instance, AKind))
}

