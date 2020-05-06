
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

type TBevel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBevel(owner IComponent) *TBevel {
    b := new(TBevel)
    b.instance = Bevel_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsBevel(obj interface{}) *TBevel {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TBevel{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsBevel.
func BevelFromInst(inst uintptr) *TBevel {
    return AsBevel(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsBevel.
func BevelFromObj(obj IObject) *TBevel {
    return AsBevel(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsBevel.
func BevelFromUnsafePointer(ptr unsafe.Pointer) *TBevel {
    return AsBevel(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (b *TBevel) Free() {
    if b.instance != 0 {
        Bevel_Free(b.instance)
        b.instance, b.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBevel) Instance() uintptr {
    return b.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBevel) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBevel) IsValid() bool {
    return b.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (b *TBevel) Is() TIs {
    return TIs(b.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (b *TBevel) As() TAs {
//    return TAs(b.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBevelClass() TClass {
    return Bevel_StaticClassType()
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (b *TBevel) BringToFront() {
    Bevel_BringToFront(b.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (b *TBevel) ClientToScreen(Point TPoint) TPoint {
    return Bevel_ClientToScreen(b.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (b *TBevel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (b *TBevel) Dragging() bool {
    return Bevel_Dragging(b.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (b *TBevel) HasParent() bool {
    return Bevel_HasParent(b.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (b *TBevel) Hide() {
    Bevel_Hide(b.instance)
}

// CN: 要求重绘。
// EN: Redraw.
func (b *TBevel) Invalidate() {
    Bevel_Invalidate(b.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (b *TBevel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Bevel_Perform(b.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (b *TBevel) Refresh() {
    Bevel_Refresh(b.instance)
}

// CN: 重绘。
// EN: Repaint.
func (b *TBevel) Repaint() {
    Bevel_Repaint(b.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (b *TBevel) ScreenToClient(Point TPoint) TPoint {
    return Bevel_ScreenToClient(b.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (b *TBevel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (b *TBevel) SendToBack() {
    Bevel_SendToBack(b.instance)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (b *TBevel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Bevel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 显示控件。
// EN: Show control.
func (b *TBevel) Show() {
    Bevel_Show(b.instance)
}

// CN: 控件更新。
// EN: Update.
func (b *TBevel) Update() {
    Bevel_Update(b.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (b *TBevel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Bevel_GetTextBuf(b.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (b *TBevel) GetTextLen() int32 {
    return Bevel_GetTextLen(b.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (b *TBevel) SetTextBuf(Buffer string) {
    Bevel_SetTextBuf(b.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (b *TBevel) FindComponent(AName string) *TComponent {
    return AsComponent(Bevel_FindComponent(b.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBevel) GetNamePath() string {
    return Bevel_GetNamePath(b.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBevel) Assign(Source IObject) {
    Bevel_Assign(b.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBevel) ClassType() TClass {
    return Bevel_ClassType(b.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBevel) ClassName() string {
    return Bevel_ClassName(b.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBevel) InstanceSize() int32 {
    return Bevel_InstanceSize(b.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBevel) InheritsFrom(AClass TClass) bool {
    return Bevel_InheritsFrom(b.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBevel) Equals(Obj IObject) bool {
    return Bevel_Equals(b.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBevel) GetHashCode() int32 {
    return Bevel_GetHashCode(b.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (b *TBevel) ToString() string {
    return Bevel_ToString(b.instance)
}

func (b *TBevel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Bevel_AnchorToNeighbour(b.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (b *TBevel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Bevel_AnchorParallel(b.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (b *TBevel) AnchorHorizontalCenterTo(ASibling IControl) {
    Bevel_AnchorHorizontalCenterTo(b.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (b *TBevel) AnchorVerticalCenterTo(ASibling IControl) {
    Bevel_AnchorVerticalCenterTo(b.instance, CheckPtr(ASibling))
}

func (b *TBevel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Bevel_AnchorAsAlign(b.instance, ATheAlign , ASpace)
}

func (b *TBevel) AnchorClient(ASpace int32) {
    Bevel_AnchorClient(b.instance, ASpace)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (b *TBevel) Align() TAlign {
    return Bevel_GetAlign(b.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (b *TBevel) SetAlign(value TAlign) {
    Bevel_SetAlign(b.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (b *TBevel) Anchors() TAnchors {
    return Bevel_GetAnchors(b.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (b *TBevel) SetAnchors(value TAnchors) {
    Bevel_SetAnchors(b.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (b *TBevel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Bevel_GetConstraints(b.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (b *TBevel) SetConstraints(value *TSizeConstraints) {
    Bevel_SetConstraints(b.instance, CheckPtr(value))
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (b *TBevel) ParentShowHint() bool {
    return Bevel_GetParentShowHint(b.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (b *TBevel) SetParentShowHint(value bool) {
    Bevel_SetParentShowHint(b.instance, value)
}

func (b *TBevel) Shape() TBevelShape {
    return Bevel_GetShape(b.instance)
}

func (b *TBevel) SetShape(value TBevelShape) {
    Bevel_SetShape(b.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (b *TBevel) ShowHint() bool {
    return Bevel_GetShowHint(b.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (b *TBevel) SetShowHint(value bool) {
    Bevel_SetShowHint(b.instance, value)
}

func (b *TBevel) Style() TBevelStyle {
    return Bevel_GetStyle(b.instance)
}

func (b *TBevel) SetStyle(value TBevelStyle) {
    Bevel_SetStyle(b.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBevel) Visible() bool {
    return Bevel_GetVisible(b.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBevel) SetVisible(value bool) {
    Bevel_SetVisible(b.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBevel) Enabled() bool {
    return Bevel_GetEnabled(b.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBevel) SetEnabled(value bool) {
    Bevel_SetEnabled(b.instance, value)
}

func (b *TBevel) Action() *TAction {
    return AsAction(Bevel_GetAction(b.instance))
}

func (b *TBevel) SetAction(value IComponent) {
    Bevel_SetAction(b.instance, CheckPtr(value))
}

func (b *TBevel) BiDiMode() TBiDiMode {
    return Bevel_GetBiDiMode(b.instance)
}

func (b *TBevel) SetBiDiMode(value TBiDiMode) {
    Bevel_SetBiDiMode(b.instance, value)
}

func (b *TBevel) BoundsRect() TRect {
    return Bevel_GetBoundsRect(b.instance)
}

func (b *TBevel) SetBoundsRect(value TRect) {
    Bevel_SetBoundsRect(b.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (b *TBevel) ClientHeight() int32 {
    return Bevel_GetClientHeight(b.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (b *TBevel) SetClientHeight(value int32) {
    Bevel_SetClientHeight(b.instance, value)
}

func (b *TBevel) ClientOrigin() TPoint {
    return Bevel_GetClientOrigin(b.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (b *TBevel) ClientRect() TRect {
    return Bevel_GetClientRect(b.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (b *TBevel) ClientWidth() int32 {
    return Bevel_GetClientWidth(b.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (b *TBevel) SetClientWidth(value int32) {
    Bevel_SetClientWidth(b.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (b *TBevel) ControlState() TControlState {
    return Bevel_GetControlState(b.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (b *TBevel) SetControlState(value TControlState) {
    Bevel_SetControlState(b.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (b *TBevel) ControlStyle() TControlStyle {
    return Bevel_GetControlStyle(b.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (b *TBevel) SetControlStyle(value TControlStyle) {
    Bevel_SetControlStyle(b.instance, value)
}

func (b *TBevel) Floating() bool {
    return Bevel_GetFloating(b.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBevel) Parent() *TWinControl {
    return AsWinControl(Bevel_GetParent(b.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBevel) SetParent(value IWinControl) {
    Bevel_SetParent(b.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (b *TBevel) Left() int32 {
    return Bevel_GetLeft(b.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (b *TBevel) SetLeft(value int32) {
    Bevel_SetLeft(b.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (b *TBevel) Top() int32 {
    return Bevel_GetTop(b.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (b *TBevel) SetTop(value int32) {
    Bevel_SetTop(b.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (b *TBevel) Width() int32 {
    return Bevel_GetWidth(b.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (b *TBevel) SetWidth(value int32) {
    Bevel_SetWidth(b.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (b *TBevel) Height() int32 {
    return Bevel_GetHeight(b.instance)
}

// CN: 设置高度。
// EN: Set height.
func (b *TBevel) SetHeight(value int32) {
    Bevel_SetHeight(b.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBevel) Cursor() TCursor {
    return Bevel_GetCursor(b.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBevel) SetCursor(value TCursor) {
    Bevel_SetCursor(b.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (b *TBevel) Hint() string {
    return Bevel_GetHint(b.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (b *TBevel) SetHint(value string) {
    Bevel_SetHint(b.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBevel) ComponentCount() int32 {
    return Bevel_GetComponentCount(b.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (b *TBevel) ComponentIndex() int32 {
    return Bevel_GetComponentIndex(b.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (b *TBevel) SetComponentIndex(value int32) {
    Bevel_SetComponentIndex(b.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBevel) Owner() *TComponent {
    return AsComponent(Bevel_GetOwner(b.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBevel) Name() string {
    return Bevel_GetName(b.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBevel) SetName(value string) {
    Bevel_SetName(b.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBevel) Tag() int {
    return Bevel_GetTag(b.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBevel) SetTag(value int) {
    Bevel_SetTag(b.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (b *TBevel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideLeft(b.instance))
}

// CN: 设置左边锚点。
// EN: .
func (b *TBevel) SetAnchorSideLeft(value *TAnchorSide) {
    Bevel_SetAnchorSideLeft(b.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (b *TBevel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideTop(b.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (b *TBevel) SetAnchorSideTop(value *TAnchorSide) {
    Bevel_SetAnchorSideTop(b.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (b *TBevel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideRight(b.instance))
}

// CN: 设置右边锚点。
// EN: .
func (b *TBevel) SetAnchorSideRight(value *TAnchorSide) {
    Bevel_SetAnchorSideRight(b.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (b *TBevel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideBottom(b.instance))
}

// CN: 设置底边锚点。
// EN: .
func (b *TBevel) SetAnchorSideBottom(value *TAnchorSide) {
    Bevel_SetAnchorSideBottom(b.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (b *TBevel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Bevel_GetBorderSpacing(b.instance))
}

// CN: 设置边框间距。
// EN: .
func (b *TBevel) SetBorderSpacing(value *TControlBorderSpacing) {
    Bevel_SetBorderSpacing(b.instance, CheckPtr(value))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBevel) Components(AIndex int32) *TComponent {
    return AsComponent(Bevel_GetComponents(b.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (b *TBevel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSide(b.instance, AKind))
}

