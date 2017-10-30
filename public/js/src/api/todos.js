/* global fetch */

export const getTodos = () =>
  fetch('/todos')
    .then(response => response.json());

export const addTodo = () => {};
