<template>
  <div>
    <h1>
      todo list app
    </h1>
    <hr>
    <h2>
      Todos
    </h2>
    <todo-list></todo-list>
    <add-todo></add-todo>
  </div>
</template>

<script>
import TodoList from './TodoList.vue';
import AddTodo from './AddTodo.vue';

export default {
  components: { TodoList, AddTodo }
}
</script>

