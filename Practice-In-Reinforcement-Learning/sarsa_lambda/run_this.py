from sarsa_lambda.maze_env import Maze
from sarsa_lambda.RL_brain import SarsaLambda


def update():
    for episode in range(100):
        observation = env.reset()
        action = RL.choose_action(str(observation))
        RL.eligibility_trace *= 0

        while True:
            # fresh env
            observation_, reward, done = env.step(action)
            env.render()

            action_ = RL.choose_action(str(observation_))
            RL.learn(str(observation), action, reward, str(observation_), action_)

            observation = observation_
            action = action_

            if done:
                env.render()
                break
    # end of game
    print('game over')
    env.destroy()


if __name__ == '__main__':
    env = Maze()
    RL = SarsaLambda(actions=list(range(env.n_actions)))

    env.after(100, update)
    env.mainloop()
