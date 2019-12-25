from q_learning_maze.maze_env import Maze
from q_learning_maze.RL_brain import QLearningTable


def update():
    for episode in range(100):
        observation = env.reset()

        while True:
            # fresh env
            action = RL.choose_action(str(observation))

            observation_, reward, done = env.step(action)
            env.render()

            RL.learn(str(observation), action, reward, str(observation_))

            observation = observation_

            if done:
                env.render()
                break
    # end of game
    print('game over')
    env.destroy()


if __name__ == '__main__':
    env = Maze()
    RL = QLearningTable(actions=list(range(env.n_actions)))

    env.after(100, update)
    env.mainloop()
