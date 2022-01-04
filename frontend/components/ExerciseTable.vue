<template>
  <div class="container-fluid">
    <div>
      <b-button variant="primary"
        ><NuxtLink to="./create">Create new Exercise</NuxtLink></b-button
      >
    </div>
    <b-table striped hover :items="exercises" :fields="fields"></b-table>
  </div>
</template>

<script>
export default {
  data: () => ({
    fields: ["exercise_name", "description", "target_muscle"],
    exercises: [],
  }),
  async fetch() {
    this.exercises = await fetch("http://localhost:5000/exercises").then(
      (res) => res.json()
    );
  },
};
</script>

<style>
.container {
  margin: 0.3em;
  padding: 0.1em;
}
</style>
