import * as todosServer from '../../api/todos';
import * as types from '../mutation-types';

const defaultState = {
  todos: [{ id: 1, title: 'test1', isCompleted: true }],
};

const actions = {
  getAllTodos({ commit }) {
    todosServer.getTodos()
      .then((todos) => {
        console.log('todos', todos);
        const remoteTodos = todos.map(todo => ({ ...todo, id: Date.now() }));
        commit(types.RECEIVE_TODOS, remoteTodos);
      });
  },
  toggleTodo({ commit }, todo) {
    commit(types.TOGGLE_TODO, todo);
  },
  addTodo({ commit }, text) {
    commit(types.ADD_TODO, text);
  },
};

const getters = {
  allTodos: state => state.todos,
};

const mutations = {
  [types.TOGGLE_TODO](state, { id, isCompleted }) {
    const todoIndex = state.todos.findIndex(t => t.id === id);
    state.todos[todoIndex].isCompleted = !isCompleted;
  },
  [types.ADD_TODO](state, text) {
    const newTodo = { id: Date.now(), title: text, isCompleted: false };
    state.todos.push(newTodo);
  },
};

export default {
  actions,
  state: defaultState,
  getters,
  mutations,
};
