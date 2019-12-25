## 胸部暴露识别
采用迁移学习，代码大多参考《tensorflow实战》。将Inception的最后一个池化层输出作为bottleneck，
后面单独接了一个softmax，没有做finetune。

## 数据集
不可能有的，兄嘚

## inception模型下载
`get_inception_v3_model.py`

## 测试inception模型
`get_inception_prediction.py`

## 数据集（以类别作为文件名）
/data/train/positive/  存放正样本图片
/data/train/negative/  存放负样本图片

## 数据增强
`image_augment.py` 创造一些图片
需要手动将创造出的图片移动到正负样本的文件夹内

## 模型训练
上面这些处理完后，可以训练了。
`transfer_model.py`
注释非常详细，按需修改
其中，`loss_visualization.py`记录了train loss和validation loss的学习曲线以及最后几千步的loss。
**LR模型的以npz的形式存在/lr_model/下。**

## 模型部署
`transfer_serving.py`
基本上500ms(cpu机器)左右预测一张图片，采用多进程，将敏感id写入mysql。

## 总结
真实测试发现，预测的分数达到0.95以上，都是真正的暴露凶器。Inception模型的特征提取作用非常靠谱。
1. 数据集很重要，但我不能提供。
1. 二分类其实不太准确，应该是多分类（暴露，非暴露，其他），但是打标签太麻烦了。这导致分数中间的图片不太明确。
1. 可以添加正则项。
1. 修改下输入，就可以部署预测一些图片，0.95以上的请分享
