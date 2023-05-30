# 作业3要求

参考《代码中的软件工程》第六章可复用软件设计及lab5.2的源代码，完成实验并写一篇实验报告，总结Callback函数的工作机制以及通过参数进行解耦合的方法，深入理解接口设计中的抽象分层。



# lab5.2的文件结构

通用的LinkTable模块的接口在linktable.h，对应的代码实现在linktable.c，linktableinternal.h存放了LinkTable模块的结构体定义。

linktableinternal.h：定义了结构体LinkTableNode和结构体LinkTable。LinkTableNode是LinkTable链表中的一个节点，成员变量是指向下一个节点的指针。LinkTable是一个链表，成员变量是头节点、尾节点和节点数量，以及一个互斥变量mutex。

linktable.h：存放了通用的LinkTable模块的接口，将LinkTableNode和LinkTable重命名为tLinkTableNode何tLingkTable，并基于这两个结构体定义了一系列接口。提供的接口有：创建链表、删除链表、向链表中增加节点、从链表中删除节点、从链表中搜索节点、获取头节点、获取下一节点的操作。

linktable.c：存放了linktable.h中定义的接口的代码实现。

menu.c：实现menu功能的文件，使用linktable.h定义的接口，实现最终功能。



# 回调函数（callback）的定义

回调函数是指一段以参数的形式传递给其他代码的可执行代码，其思想可以比较通俗地表达为：“程序员知道应该做什么、但我们不知道什么时候去做，只有其他模块知道什么时候去做，因此程序员需要把应该做什么封装成回调函数告诉其他模块”。

调用（callin）就是程序员调用预先写好的系统的函数；而回调（callback）就是程序员写一个函数，让预先写好的系统来调用。



# 以menu中的例子来进一步分析回调

LinkTable模块的SearchLinkTableNode方法使用了callback接口，其代码实现如下。SearchLinkTableNode函数是callin函数，而Condition函数是callin函数的参数，Condition函数就是callback函数。

```c
tLinkTableNode * SearchLinkTableNode(tLinkTable *pLinkTable, 
                        int Condition(tLinkTableNode * pNode, void * args),
                        void * args)
{
    if(pLinkTable == NULL || Condition == NULL)
    {
        return NULL;
    }
    tLinkTableNode * pNode = pLinkTable->pHead;
    while(pNode != NULL)
    {    
        if(Condition(pNode, args) == SUCCESS)
        {
            return pNode;				    
        }
        pNode = pNode->pNext;
    }
    return NULL;
}
```

我们还可以注意到，callin函数增加了一个参数args，callback函数同样增加了一个参数args函数。这里我们结合menu.c的使用情景来说明。

SearchCondition函数定义了查询条件，说明了什么样的节点是符合条件的，这个方法可以实现用户自定义查询条件，可以和应用背景结合。FindCmd函数的目标是找到一个对应的cmd并返回节点指针。在callin函数和callback函数增加参数args的原因是尽可能追求松散耦合，args用于传递用户输入的菜单命令（help、version或quit）。这样做的原因是，可以使查询接口更加通用，有效地提高了接口的通用性，同时尽可能使得各模块之间松散耦合。

```c
int SearchCondition(tLinkTableNode * pLinkTableNode, void * args)
{
    char * cmd = (char*) args;
    tDataNode * pNode = (tDataNode *)pLinkTableNode;
    if(strcmp(pNode->cmd, cmd) == 0)
    {
        return  SUCCESS;  
    }
    return FAILURE;	       
}

/* find a cmd in the linklist and return the datanode pointer */
tDataNode* FindCmd(tLinkTable * head, char * cmd)
{
    return  (tDataNode*)SearchLinkTableNode(head, SearchCondition, (void*)cmd);
}
```

在这里补充一下几种耦合类型的定义：

耦合度是指软件模块之间的依赖程度，一般可以分为紧密耦合（Tightly Coupled）、松散耦合（Loosely Coupled）和无耦合（Uncoupled），在一般的软件设计中我们追求松散耦合。

更细致地对耦合度进行进一步划分的话，耦合度依此递增可以分为无耦合、数据耦合、标记耦合、控制耦合、公共耦合和内容耦合。这些耦合度的划分依据就是接口的定义方式。下面给出了数据耦合、标记耦合和公共耦合的定义：

1、数据耦合：在软件模块之间仅通过显式的调用传递基本数据类型即为数据耦合。

2、标记耦合：在软件模块之间仅通过显式的调用传递复杂的数据结构(结构化数据）即为标记耦合，这时数据的结构成为调用双方软件模块隐含的规格约定，因此耦合度要比数据耦合高。但相比公共耦合没有经过显式的调用传递数据的方式耦合度要低。

3、公共耦合：当软件模块之间共享数据区或变量名的软件模块之间即是公共耦合，显然两个软件模块之间的接口定义不是通过显式的调用方式，而是隐式的共享了共享了数据区或变量名。

上面提到的在callin函数和callback函数中增加args参数，在各个模块之间通过显示地传递复杂数据结构（一个cmd数组），属于标记耦合。

最后概括一下这个查询功能体现的接口设计分层思想：

1、linktableinternal.h：定义了链表节点和链表的基本结构

2、linktable.h：定义了模块接口，在这里定义了callin函数

3、linktable.c：实现了linktable.h中定义的模块接口，定义callin函数如何调用callback函数并如何做出反应

4、menu.c：根据用户需求自定义查询条件（定义callback函数），调用callin函数





作者：093