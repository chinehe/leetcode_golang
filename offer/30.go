package offer

// MinStackWithSlice 使用Slice实现双栈
type MinStackWithSlice struct {
	data       []int // 数据栈
	sortedData []int // 数据栈的非严格降序栈
}

// MinStackWithSliceConstructor initialize your data structure here.
func MinStackWithSliceConstructor() MinStackWithSlice {
	return MinStackWithSlice{
		data:       make([]int, 0),
		sortedData: make([]int, 0),
	}
}

func (this *MinStackWithSlice) Push(x int) {
	// 数据栈直接入
	this.data = append(this.data, x)
	// 如果辅助栈为空或，栈顶元素大于x，入栈
	size := len(this.sortedData)
	if size == 0 || this.sortedData[size-1] > x {
		this.sortedData = append(this.sortedData, x)
	}
}

func (this *MinStackWithSlice) Pop() {
	// 获取数据栈栈顶元素并移除
	dataSize := len(this.data)
	if dataSize == 0 {
		return
	}
	data := this.data[dataSize-1]
	this.data = this.data[0 : dataSize-1]
	// 如果辅助栈的栈顶元素等于当前元素，也出栈
	sortedSize := len(this.sortedData)
	if sortedSize > 0 && this.sortedData[sortedSize-1] == data {
		this.sortedData = this.sortedData[0 : sortedSize-1]
	}
}

func (this *MinStackWithSlice) Top() int {
	dataSize := len(this.data)
	if dataSize == 0 {
		return -1
	}
	return this.data[dataSize-1]
}

func (this *MinStackWithSlice) Min() int {
	sortedSize := len(this.sortedData)
	if sortedSize == 0 {
		return -1
	} else {
		return this.sortedData[sortedSize-1]
	}
}
