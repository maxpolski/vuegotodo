/* global fetch */

export const getTodos = () =>
  fetch('/todos')
    .then(response => response.json());

export const addTodo = () => {};

export const toggleTodo = todoText =>
  fetch('/todos', {
    method: 'PUT',
    body: JSON.stringify({ todoText }),
  }).then(resp => resp.json())
    .then(data => data);
