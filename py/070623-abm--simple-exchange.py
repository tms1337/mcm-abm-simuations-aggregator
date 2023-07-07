# ===> sampkle

import numpy as np
import scipy
import seaborn as sns
import matplotlib.pyplot as plt

def apply_limit_buy(price, amount, state):
    state.buy_orders.append(("LIMIT_BUY", price, amount))
    state.buy_orders.sort(reverse=True)

def apply_limit_sell(price, amount, state):
    state.sell_orders.append(("LIMIT_SELL", price, amount))
    state.sell_orders.sort()

def apply_market_buy(amount, state):
    selected_orders = []
    while total_amount(selected_orders) < amount:
        if len(state.sell_orders) == 0:
            state.sell_orders = []
            state.price = 0
            break
        selected_orders.append(state.sell_orders.pop(0))
    n = len(selected_orders)
    if n >= len(state.sell_orders):
        state.sell_orders = []
        state.price = 0
    else:
        state.sell_orders = state.sell_orders[n:]
        state.price = state.sell_orders[0][1]

def apply_market_sell(amount, state):
    selected_orders = []
    while total_amount(selected_orders) < amount:
        if len(state.buy_orders) == 0:
            state.buy_orders = []
            state.price = 0
            break
        selected_orders.append(state.buy_orders.pop(0))
    n = len(selected_orders)
    if n >= len(state.buy_orders):
        state.buy_orders = []
        state.price = 0
    else:
        state.buy_orders = state.buy_orders[n:]
        state.price = state.buy_orders[0][1]

def total_amount(orders):
    return sum(order[2] for order in orders)

def plot_price_history(price_history):
    # sns.lineplot(range(len(price_history)), price_history)
    plt.plot(price_history)
    plt.show()

def initialize_constants():
    total_steps = 1000
    total_agents = 50
    return total_steps, total_agents

def initialize_world():
    state = ddict({"price": 1.0, "buy_orders": [], "sell_orders": []})
    price_history = []
    return state, price_history

def generate_random_action(state):
    action_type = np.random.choice(["LIMIT_BUY", "LIMIT_SELL", "MARKET_BUY", "MARKET_SELL"])
    if action_type in ["LIMIT_BUY", "LIMIT_SELL"]:
        price = np.random.normal(state["price"], state["price"] * 0.1)
        amount = np.random.randint(1, 10)
        return (action_type, price, amount)
    else:
        amount = np.random.randint(1, 10)
        return (action_type, amount)

def apply_actions(step_actions, state):
    for action in step_actions:
        if action[0] == "LIMIT_BUY":
            apply_limit_buy(action[1], action[2], state)
        elif action[0] == "LIMIT_SELL":
            apply_limit_sell(action[1], action[2], state)
        elif action[0] == "MARKET_BUY":
            apply_market_buy(action[1], state)
        elif action[0] == "MARKET_SELL":
            apply_market_sell(action[1], state)

def simulate_trading():
    total_steps, total_agents = initialize_constants()
    state, price_history = initialize_world()

    for step in range(total_steps):
        step_actions = []
        total_actions = np.random.randint(0, total_agents)
        for _ in range(total_actions):
            action = generate_random_action(state)
            step_actions.append(action)
        apply_actions(step_actions, state)
        price_history.append(state["price"])

    plot_price_history(price_history)

simulate_trading()
