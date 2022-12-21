# Data Structures
## Array
连续的存储空间存放数据元素
- [ ] implement
    - [ ] Get(n)    -- return the value of the nth elements 
    - [ ] Insert(n, v)  -- insert a element at given index
    - [ ] Remove(n)     -- return the removed element at given index
    - [ ] RemoveValue(v)     -- return the first element in the list with this value
    - [ ] Size()  -- return number of data elements in list
    - [ ] Capacity()  -- return number of elements it can hold
## List
任意的存储空间串起来存放数据元素。
为什么用链表
- 解决数量规模大时，数组这种连续存储结构动态维护成本高
- 预分配存储空间不好估算，分大了，浪费空间，分笑了，扩容浪费性能
链表存储基本单元结点，由两部分构成，
- 存储数据元素信息称为数据域
- 存储后续结点位置称为指针域
### Linked List
- [ ] implement
    - [ ] Insert(index, value) - insert value at index, so current item at that index is pointed to by new item at index
    - [ ] PushFront(value) - adds an item to the front of the list
    - [ ] PopFront() - remove front item and return its value
    - [ ] PushBack(value) - adds an item at the end
    - [ ] PopBack() - removes end item and returns its value
    - [ ] Size() - returns number of data elements in list
    - [ ] IsEmpty() - bool returns true if empty
    - [ ] ValueAtIndex(index) - returns the value of the nth item (starting at 0 for first)
    - [ ] Reverse() - reverses the list
    - [ ] RemoveValue(value) - removes the first item in the list with this value
### 静态链表
静态链表是一种数组描述的链表，数组的的元素都是由两个数据域组成：`data`和`cur`。
`data`: 用于存放数据；`cur`: 用于存放下一个数据在数组中的下标（相当与链表中的`next`指针）
- [ ] implement
    - [ ] Insert(index, value)
    - [ ] Remove(index, value)
### 循环链表
单链表的尾部结点的指针指向头结点
### 双向链表
链表存储基本单元结点，由三部分构成，
- 存储前面一个结点位置称为指针域
- 存储数据元素信息称为数据域
- 存储后续结点位置称为指针域

## 栈 Queue
### Tree