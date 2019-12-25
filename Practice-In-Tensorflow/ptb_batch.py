import numpy as np
import tensorflow as tf

TRAIN_DATA = ""
TRAIN_BATCH_SZIE = 20
TRAIN_NUM_STEP = 35


def read_data(file_path):
    with open(file_path, "r") as f:
        id_string = ' '.join([line.strip() for line in f.readlines()])
        return id_string


def make_batches(id_list, batch_size, num_step):
    # calculate num of batches
    num_batches = (len(id_list) -1)/(batch_size * num_step)
    data = np.array(id_list[:num_batches*batch_size*num_step])
    data = np.reshape(data, [batch_size, num_batches * num_step])
    data_batches = np.split(data, num_batches, axis=1)

    # repeat above operations and right shift a word id.
    # they are the labels
    label = np.array(id_list[1:num_batches*batch_size*num_step+1])
    label = np.reshape(label, [batch_size, num_batches * num_step])
    label_batches = np.split(label, num_batches,axis=1)
    return list(zip(data_batches, label_batches))


def main():
    train_batches = make_batches(read_data(TRAIN_DATA), TRAIN_BATCH_SZIE, TRAIN_NUM_STEP)
    pass
