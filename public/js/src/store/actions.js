import * as types from './mutation-types';

export const toggleTodo = ({ commit }, todo) => {
  commit(types.TOGGLE_TODO, {
    todo,
  })
};

