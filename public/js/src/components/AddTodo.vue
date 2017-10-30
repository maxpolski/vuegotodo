<template>
  <form>
    <input type="text" v-model="newTodo">
    <button @click="addTodo">Add</button>
  </form>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  data: () => ({
    newTodo: '',
  }),
  methods: {
    addTodo(e) {
      e.preventDefault();
      if (this.newTodo !== '') {
        this.$store.dispatch('addTodo', this.newTodo);
        this.newTodo = '';
      }
    }
  },
}
</script>
