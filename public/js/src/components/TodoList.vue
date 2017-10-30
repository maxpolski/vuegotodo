<template>
  <ul>
    <li
      v-for="t in todos" :key="t.id" :class="{completed : t.isCompleted}"
      @click="toggleTodo(t)"
    >
      {{ t.title }}
    </li>
  </ul>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  created() {
    this.$store.dispatch('getAllTodos');
  },
  methods: mapActions([
    'toggleTodo',
  ]),
  computed: mapGetters({
    todos: 'allTodos',
  })
}
</script>
