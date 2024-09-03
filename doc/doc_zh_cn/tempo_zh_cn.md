# 节奏

此项目翻译自[tempo.md](../tempo.md)

乐谱的**节奏**描述了音符演奏的快慢程度 这个值通常以每分钟节拍数(BPM)表示 与音符的时值(如四分音符 二分音符 全音符)结合使用 以确定音符的持续时间(以毫秒为单位)

## `tempo`属性

在Alda中 指定节奏的最简单方法是通过`tempo`属性 它可以以BPM指定节奏 其中每个节拍占用一个四分音符的长度(以四分音符为一拍)

例如 要指定180BPM的节奏 (♩ = 180):

```alda
(tempo! 180)
```

在传统音乐符号中 通常会看到以每分钟节拍数描述的节奏 但是有的拍号不是以四分音符为一拍 例如 节奏可以用"每分钟二分音符"来表示 即 𝅗𝅥 = 100 这仍然是"每分钟节拍" 只是它是以二分音符为一拍的

Alda的`tempo`功能非常灵活 它允许您指定以什么时值的音符为一拍

```alda
(tempo! 2 100)
```

在三拍子中 通常是以带附点的音符为一拍 例如在6/8拍下 每个小节通常表示为2拍 以附点四分音符(♩.)为一拍 这种情况下 用作为节拍的音符时值来表示节奏就很方便

`tempo`属性的音符值参数必须是有效的数字或字符串 刚刚的例子中 `4.`不是有效数字 您需要用一对双引号将其声明为字符串

```alda
# ♩. = 100
(tempo! "4." 100)
```

## 节拍转换

您也可能通过[节拍转换](https://en.wikipedia.org/wiki/Metric_modulation)来表达节奏 即从一个节拍转换到另一个节拍

例如 假设您在编写一个以9/8开头的乐谱--以附点四分音符为一拍 每小节3拍

在乐曲中 您想要过渡到3/2拍--每小节仍然是3拍 但是现在以二分音符为一拍 您希望曲速保持不变(一拍的实际时长不变) 将现在每个节拍细分为4个八分音符 而不是3个 要怎么做呢

在传统记谱法中 当乐谱的拍号发生变化时 通常会有"♩. = 𝅗𝅥 "这样的注释 在那之后 拍子保持不变 但过去表示附点四分音符的时间量现在表示二分音符 当管弦乐队演奏到乐谱中的那个地方时 指挥家继续以相同的"速度"指挥 但每位音乐家都会在心里调整自己跟拍子的思路 在心理上将每拍看作四个八分音符 而不是3个八分音符

在Alda中 您可以将"♩. = 𝅗𝅥 "这样的节拍转换表示为:

```alda
(metric-modulation! "4." 2)
```
