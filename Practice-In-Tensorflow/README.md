# tensorflow
copied in <Tensorflow实战Google深度学习框架>  


## LeNet5
After 1 training steps, loss on training batch is 5.94812.  
After 1001 training steps, loss on training batch is 0.776025.  
After 2001 training steps, loss on training batch is 0.776466.  

...  
...  
...  
After 27001 training steps, loss on training batch is 0.608183.  
After 28001 training steps, loss on training batch is 0.610333.  
After 29001 training steps, loss on training batch is 0.609691.   

*EVAL:*
After 29001 training steps, validation accuracy = 0.9898

## Transfer_Learning
>https://www.jianshu.com/p/cc830a6ed54b

## Image Pretreatment
![avatar](https://github.com/yinyajun/tensorflow/blob/master/lena_pretreatkment.jpg)


## Multi threading 
*Cooredinator*  
Coordinator类用来帮助多个线程协同工作，多个线程同步终止。 其主要方法有：  
should_stop():如果线程应该停止则返回True。
request_stop(<exception>): 请求该线程停止。
join(<list of threads>):等待被指定的线程终止。
首先创建一个Coordinator对象，然后建立一些使用Coordinator对象的线程。这些线程通常一直循环运行，一直到should_stop()返回True时停止。 任何线程都可以决定计算什么时候应该停止。它只需要调用request_stop()，同时其他线程的should_stop()将会返回True，然后都停下来。


## LSTM
![avatar](https://github.com/yinyajun/tensorflow/blob/master/lstm_predict_sin.png)
