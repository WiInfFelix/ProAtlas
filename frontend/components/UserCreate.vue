<template>
  <div>
    <b-form id="createForm" @submit="onSubmit">
      <b-form-group
        id="input-group-1"
        label="Username"
        label-for="input-1"
        description="Choose a username"
      >
        <b-form-input
          id="input-1"
          v-model="form.username"
          type="text"
          placeholder="Username"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group
        id="input-group-2"
        label="Email"
        label-for="input-2"
        description="Provide your email"
      >
        <b-form-input
          id="input-2"
          v-model="form.email"
          type="email"
          placeholder="Email"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group
        id="input-group-3"
        label="Password"
        label-for="input-3"
        description="Choose a password"
      >
        <b-form-input
          id="input-3"
          v-model="form.password"
          type="password"
          placeholder="Password"
          required
        ></b-form-input>
      </b-form-group>
      <b-button type="submit" variant="primary">Create user</b-button>
    </b-form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      form: {
        username: "",
        email: "",
        password: "",
      },
    };
  },
  methods: {
    onSubmit(event) {
      event.preventDefault();
      axios
        .post("http://localhost:5000/users", this.form)
        .then((res) => {
          alert("User created!");
        })
        .catch((error) => {
          alert(error);
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
