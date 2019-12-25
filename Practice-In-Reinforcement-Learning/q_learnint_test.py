import pandas as pd
import numpy as np
import time

np.random.seed(3)

N_STATES = 6
ACTIONS = ['left', 'right']
EPSILON = 0.9
ALPHA = 0.1
GAMMA = 0.9
MAX_EPISODES = 15  # 最大回合数
FRESH_TIME = 0.1  # 移动间隔时间


def init_q_table(n_states, actions):
    table = pd.DataFrame(
        np.zeros([n_states, len(actions)], dtype=np.float32),
        columns=actions)
    # print(table)
    return table


def choose_action(state, q_table):
    state_actions = q_table.iloc[state, :]
    if np.random.uniform() > EPSILON or state_actions.all() == 0:
        action = np.random.choice(ACTIONS)
    else:
        action = state_actions.argmax()  # Series的argmax返回的是index
    return action


def get_env_feedback(state, action):
    if action == 'right':
        if state == N_STATES - 2:
            R = 1
            S_ = 'terminal'
        else:
            R = 0
            S_ = state + 1
    else:
        if state == 0:
            S_ = state
            R = 0
        else:
            R = 0
            S_ = state - 1
    return S_, R


def update_env(S, episode, step_counter):
    env_list = ['_'] * (N_STATES - 1) + ['T']
    if S == 'terminal':
        interaction = 'Episode %s: total_steps = %s' % (episode + 1, step_counter)
        print('\r{}'.format(interaction), end='')
        time.sleep(1)
        print('\r                                ', end='')
    else:
        env_list[S] = 'o'
        interaction = ''.join(env_list)
        print('\r{}'.format(interaction), end='')
        time.sleep(FRESH_TIME)


def rl():
    q_table = init_q_table(N_STATES, ACTIONS)
    for episode in range(MAX_EPISODES):
        step_counter = 0
        S = 0
        is_terminal = False
        update_env(S, episode, step_counter)
        while not is_terminal:
            A = choose_action(S, q_table)
            S_, R = get_env_feedback(S, A)
            if S_ == 'terminal':
                is_terminal = True
                q_target = R
            else:
                q_target = R + GAMMA * (q_table.iloc[S_, :].max())
            # update q_table
            q_predict = q_table.ix[S, A]
            q_table.ix[S, A] += ALPHA * (q_target - q_predict)
            S = S_

            update_env(S, episode, step_counter + 1)
            step_counter += 1
    return q_table


if __name__ == '__main__':
    q_table = rl()
    print('\rQ_table:\n')
    print(q_table)
