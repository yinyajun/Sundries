# import jieba
#
# with open("data/en-zh/train.tags.en-zh.en", 'r', encoding='utf-8') as fin:
#     for line in fin:
#         separated_line = ' '.join(jieba.cut(line))
#         with open('train.txt.en', 'a', encoding='utf-8') as fout:
#             fout.write(separated_line)
import tensorflow as tf

MAX_LEN = 50
SOS_ID = 1


def make_dataset(file_path):
    dataset = tf.data.TextLineDataset(file_path)
    dataset = dataset.map(lambda string: tf.string_split([string]).values)
    dataset = dataset.map(lambda string: tf.string_to_number(string, tf.int32))
    dataset = dataset.map(lambda x: (x, tf.size(x)))

    return dataset


def make_src_trg_dataset(src_path, trg_path, batch_size):
    src_data = make_dataset(src_path)
    trg_data = make_dataset(trg_path)
    datasset = tf.data.Dataset.zip((src_data, trg_data))
