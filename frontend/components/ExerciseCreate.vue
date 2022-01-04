<template>
  <div>
    <b-form id="createForm" @submit="onSubmit">
      <b-form-group
        id="input-group-1"
        label="Exercise name"
        label-for="input-1"
        description="Name the exercise"
      >
        <b-form-input
          id="input-1"
          v-model="form.exercise_name"
          type="text"
          placeholder="Exercise name"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group
        id="input-group-2"
        label="Description"
        label-for="input-2"
        description="Describe the exercise in short"
      >
        <b-form-textarea
          id="input-2"
          v-model="form.description"
        ></b-form-textarea>
      </b-form-group>
      <b-form-group
        id="input-group-3"
        label="Target muscle"
        label-for="input-3"
        description="Name the target muscle"
      >
        <b-form-input
          id="input-3"
          v-model="form.target_muscle"
          type="text"
          placeholder="Muscle name"
          required
        ></b-form-input>
      </b-form-group>
      <b-button type="submit" variant="primary">Save it!!</b-button>
    </b-form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      form: {
        exercise_name: "",
        description: "",
        target_muscle: "",
      },
    };
  },
  methods: {
    onSubmit(event) {
      event.preventDefault();
      axios
        .post("http://localhost:5000/exercises", this.form)
        .then((res) => {
          alert("Exercise created!");
        })
        .catch((error) => {
          alert("There was an error creating the exercise...");
        });
    },
  },
};
</script>

<style>
#createForm {
  margin: 1em;
}
</style>
