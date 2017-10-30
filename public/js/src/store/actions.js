import * as types from './mutation-types';

export const load = ({ commit }, todo) => {
  commit(types.TOGGLE_TODO, {
    todo,
  })
};

