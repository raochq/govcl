
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

type TGauge struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGauge(owner IComponent) *TGauge {
    g := new(TGauge)
    g.instance = Gauge_Create(CheckPtr(owner))
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsGauge(obj interface{}) *TGauge {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TGauge{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsGauge.
func GaugeFromInst(inst uintptr) *TGauge {
    return AsGauge(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsGauge.
func GaugeFromObj(obj IObject) *TGauge {
    return AsGauge(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsGauge.
func GaugeFromUnsafePointer(ptr unsafe.Pointer) *TGauge {
    return AsGauge(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (g *TGauge) Free() {
    if g.instance != 0 {
        Gauge_Free(g.instance)
        g.instance, g.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGauge) Instance() uintptr {
    return g.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGauge) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGauge) IsValid() bool {
    return g.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (g *TGauge) Is() TIs {
    return TIs(g.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (g *TGauge) As() TAs {
//    return TAs(g.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGaugeClass() TClass {
    return Gauge_StaticClassType()
}

func (g *TGauge) AddProgress(Value int32) {
    Gauge_AddProgress(g.instance, Value)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (g *TGauge) BringToFront() {
    Gauge_BringToFront(g.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (g *TGauge) ClientToScreen(Point TPoint) TPoint {
    return Gauge_ClientToScreen(g.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (g *TGauge) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ClientToParent(g.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (g *TGauge) Dragging() bool {
    return Gauge_Dragging(g.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (g *TGauge) HasParent() bool {
    return Gauge_HasParent(g.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (g *TGauge) Hide() {
    Gauge_Hide(g.instance)
}

// CN: 要求重绘。
// EN: Redraw.
func (g *TGauge) Invalidate() {
    Gauge_Invalidate(g.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (g *TGauge) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Gauge_Perform(g.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (g *TGauge) Refresh() {
    Gauge_Refresh(g.instance)
}

// CN: 重绘。
// EN: Repaint.
func (g *TGauge) Repaint() {
    Gauge_Repaint(g.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (g *TGauge) ScreenToClient(Point TPoint) TPoint {
    return Gauge_ScreenToClient(g.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (g *TGauge) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ParentToClient(g.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (g *TGauge) SendToBack() {
    Gauge_SendToBack(g.instance)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (g *TGauge) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Gauge_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 显示控件。
// EN: Show control.
func (g *TGauge) Show() {
    Gauge_Show(g.instance)
}

// CN: 控件更新。
// EN: Update.
func (g *TGauge) Update() {
    Gauge_Update(g.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (g *TGauge) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Gauge_GetTextBuf(g.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (g *TGauge) GetTextLen() int32 {
    return Gauge_GetTextLen(g.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (g *TGauge) SetTextBuf(Buffer string) {
    Gauge_SetTextBuf(g.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (g *TGauge) FindComponent(AName string) *TComponent {
    return AsComponent(Gauge_FindComponent(g.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (g *TGauge) GetNamePath() string {
    return Gauge_GetNamePath(g.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (g *TGauge) Assign(Source IObject) {
    Gauge_Assign(g.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGauge) ClassType() TClass {
    return Gauge_ClassType(g.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGauge) ClassName() string {
    return Gauge_ClassName(g.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGauge) InstanceSize() int32 {
    return Gauge_InstanceSize(g.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGauge) InheritsFrom(AClass TClass) bool {
    return Gauge_InheritsFrom(g.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGauge) Equals(Obj IObject) bool {
    return Gauge_Equals(g.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGauge) GetHashCode() int32 {
    return Gauge_GetHashCode(g.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (g *TGauge) ToString() string {
    return Gauge_ToString(g.instance)
}

func (g *TGauge) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Gauge_AnchorToNeighbour(g.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (g *TGauge) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Gauge_AnchorParallel(g.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (g *TGauge) AnchorHorizontalCenterTo(ASibling IControl) {
    Gauge_AnchorHorizontalCenterTo(g.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (g *TGauge) AnchorVerticalCenterTo(ASibling IControl) {
    Gauge_AnchorVerticalCenterTo(g.instance, CheckPtr(ASibling))
}

func (g *TGauge) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Gauge_AnchorAsAlign(g.instance, ATheAlign , ASpace)
}

func (g *TGauge) AnchorClient(ASpace int32) {
    Gauge_AnchorClient(g.instance, ASpace)
}

func (g *TGauge) PercentDone() int32 {
    return Gauge_GetPercentDone(g.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (g *TGauge) Align() TAlign {
    return Gauge_GetAlign(g.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (g *TGauge) SetAlign(value TAlign) {
    Gauge_SetAlign(g.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (g *TGauge) Anchors() TAnchors {
    return Gauge_GetAnchors(g.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (g *TGauge) SetAnchors(value TAnchors) {
    Gauge_SetAnchors(g.instance, value)
}

func (g *TGauge) BackColor() TColor {
    return Gauge_GetBackColor(g.instance)
}

func (g *TGauge) SetBackColor(value TColor) {
    Gauge_SetBackColor(g.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (g *TGauge) BorderStyle() TBorderStyle {
    return Gauge_GetBorderStyle(g.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (g *TGauge) SetBorderStyle(value TBorderStyle) {
    Gauge_SetBorderStyle(g.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (g *TGauge) Color() TColor {
    return Gauge_GetColor(g.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (g *TGauge) SetColor(value TColor) {
    Gauge_SetColor(g.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (g *TGauge) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Gauge_GetConstraints(g.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (g *TGauge) SetConstraints(value *TSizeConstraints) {
    Gauge_SetConstraints(g.instance, CheckPtr(value))
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (g *TGauge) Enabled() bool {
    return Gauge_GetEnabled(g.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (g *TGauge) SetEnabled(value bool) {
    Gauge_SetEnabled(g.instance, value)
}

func (g *TGauge) ForeColor() TColor {
    return Gauge_GetForeColor(g.instance)
}

func (g *TGauge) SetForeColor(value TColor) {
    Gauge_SetForeColor(g.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (g *TGauge) Font() *TFont {
    return AsFont(Gauge_GetFont(g.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (g *TGauge) SetFont(value *TFont) {
    Gauge_SetFont(g.instance, CheckPtr(value))
}

func (g *TGauge) MinValue() int32 {
    return Gauge_GetMinValue(g.instance)
}

func (g *TGauge) SetMinValue(value int32) {
    Gauge_SetMinValue(g.instance, value)
}

func (g *TGauge) MaxValue() int32 {
    return Gauge_GetMaxValue(g.instance)
}

func (g *TGauge) SetMaxValue(value int32) {
    Gauge_SetMaxValue(g.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (g *TGauge) ParentColor() bool {
    return Gauge_GetParentColor(g.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (g *TGauge) SetParentColor(value bool) {
    Gauge_SetParentColor(g.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (g *TGauge) ParentFont() bool {
    return Gauge_GetParentFont(g.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (g *TGauge) SetParentFont(value bool) {
    Gauge_SetParentFont(g.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (g *TGauge) ParentShowHint() bool {
    return Gauge_GetParentShowHint(g.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (g *TGauge) SetParentShowHint(value bool) {
    Gauge_SetParentShowHint(g.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (g *TGauge) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Gauge_GetPopupMenu(g.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (g *TGauge) SetPopupMenu(value IComponent) {
    Gauge_SetPopupMenu(g.instance, CheckPtr(value))
}

func (g *TGauge) Progress() int32 {
    return Gauge_GetProgress(g.instance)
}

func (g *TGauge) SetProgress(value int32) {
    Gauge_SetProgress(g.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (g *TGauge) ShowHint() bool {
    return Gauge_GetShowHint(g.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (g *TGauge) SetShowHint(value bool) {
    Gauge_SetShowHint(g.instance, value)
}

func (g *TGauge) ShowText() bool {
    return Gauge_GetShowText(g.instance)
}

func (g *TGauge) SetShowText(value bool) {
    Gauge_SetShowText(g.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (g *TGauge) Visible() bool {
    return Gauge_GetVisible(g.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (g *TGauge) SetVisible(value bool) {
    Gauge_SetVisible(g.instance, value)
}

func (g *TGauge) Action() *TAction {
    return AsAction(Gauge_GetAction(g.instance))
}

func (g *TGauge) SetAction(value IComponent) {
    Gauge_SetAction(g.instance, CheckPtr(value))
}

func (g *TGauge) BiDiMode() TBiDiMode {
    return Gauge_GetBiDiMode(g.instance)
}

func (g *TGauge) SetBiDiMode(value TBiDiMode) {
    Gauge_SetBiDiMode(g.instance, value)
}

func (g *TGauge) BoundsRect() TRect {
    return Gauge_GetBoundsRect(g.instance)
}

func (g *TGauge) SetBoundsRect(value TRect) {
    Gauge_SetBoundsRect(g.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (g *TGauge) ClientHeight() int32 {
    return Gauge_GetClientHeight(g.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (g *TGauge) SetClientHeight(value int32) {
    Gauge_SetClientHeight(g.instance, value)
}

func (g *TGauge) ClientOrigin() TPoint {
    return Gauge_GetClientOrigin(g.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (g *TGauge) ClientRect() TRect {
    return Gauge_GetClientRect(g.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (g *TGauge) ClientWidth() int32 {
    return Gauge_GetClientWidth(g.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (g *TGauge) SetClientWidth(value int32) {
    Gauge_SetClientWidth(g.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (g *TGauge) ControlState() TControlState {
    return Gauge_GetControlState(g.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (g *TGauge) SetControlState(value TControlState) {
    Gauge_SetControlState(g.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (g *TGauge) ControlStyle() TControlStyle {
    return Gauge_GetControlStyle(g.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (g *TGauge) SetControlStyle(value TControlStyle) {
    Gauge_SetControlStyle(g.instance, value)
}

func (g *TGauge) Floating() bool {
    return Gauge_GetFloating(g.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (g *TGauge) Parent() *TWinControl {
    return AsWinControl(Gauge_GetParent(g.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (g *TGauge) SetParent(value IWinControl) {
    Gauge_SetParent(g.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (g *TGauge) Left() int32 {
    return Gauge_GetLeft(g.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (g *TGauge) SetLeft(value int32) {
    Gauge_SetLeft(g.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (g *TGauge) Top() int32 {
    return Gauge_GetTop(g.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (g *TGauge) SetTop(value int32) {
    Gauge_SetTop(g.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (g *TGauge) Width() int32 {
    return Gauge_GetWidth(g.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (g *TGauge) SetWidth(value int32) {
    Gauge_SetWidth(g.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (g *TGauge) Height() int32 {
    return Gauge_GetHeight(g.instance)
}

// CN: 设置高度。
// EN: Set height.
func (g *TGauge) SetHeight(value int32) {
    Gauge_SetHeight(g.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (g *TGauge) Cursor() TCursor {
    return Gauge_GetCursor(g.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (g *TGauge) SetCursor(value TCursor) {
    Gauge_SetCursor(g.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (g *TGauge) Hint() string {
    return Gauge_GetHint(g.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (g *TGauge) SetHint(value string) {
    Gauge_SetHint(g.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (g *TGauge) ComponentCount() int32 {
    return Gauge_GetComponentCount(g.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (g *TGauge) ComponentIndex() int32 {
    return Gauge_GetComponentIndex(g.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (g *TGauge) SetComponentIndex(value int32) {
    Gauge_SetComponentIndex(g.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (g *TGauge) Owner() *TComponent {
    return AsComponent(Gauge_GetOwner(g.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (g *TGauge) Name() string {
    return Gauge_GetName(g.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (g *TGauge) SetName(value string) {
    Gauge_SetName(g.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (g *TGauge) Tag() int {
    return Gauge_GetTag(g.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (g *TGauge) SetTag(value int) {
    Gauge_SetTag(g.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (g *TGauge) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideLeft(g.instance))
}

// CN: 设置左边锚点。
// EN: .
func (g *TGauge) SetAnchorSideLeft(value *TAnchorSide) {
    Gauge_SetAnchorSideLeft(g.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (g *TGauge) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideTop(g.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (g *TGauge) SetAnchorSideTop(value *TAnchorSide) {
    Gauge_SetAnchorSideTop(g.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (g *TGauge) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideRight(g.instance))
}

// CN: 设置右边锚点。
// EN: .
func (g *TGauge) SetAnchorSideRight(value *TAnchorSide) {
    Gauge_SetAnchorSideRight(g.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (g *TGauge) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideBottom(g.instance))
}

// CN: 设置底边锚点。
// EN: .
func (g *TGauge) SetAnchorSideBottom(value *TAnchorSide) {
    Gauge_SetAnchorSideBottom(g.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (g *TGauge) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Gauge_GetBorderSpacing(g.instance))
}

// CN: 设置边框间距。
// EN: .
func (g *TGauge) SetBorderSpacing(value *TControlBorderSpacing) {
    Gauge_SetBorderSpacing(g.instance, CheckPtr(value))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (g *TGauge) Components(AIndex int32) *TComponent {
    return AsComponent(Gauge_GetComponents(g.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (g *TGauge) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSide(g.instance, AKind))
}

