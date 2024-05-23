<template>
  <div id="gradient" class="gradient"></div>
  <div class="card">
    <div class="card card-container">
      <div class="form-label">
        Вход
      </div>
      <Form @submit="handleLogin" :validation-schema="schema">
        <div class="form-group">
          <label for="phoneNumber">Телефон</label>
          <Field name="phoneNumber" type="text" class="form-control" />
          <ErrorMessage name="phoneNumber" class="error-feedback" />
        </div>
        <div class="form-group">
          <label for="password">Пароль</label>
          <Field name="password" type="password" class="form-control" />
          <ErrorMessage name="password" class="error-feedback" />
        </div>

        <div class="form-group">
          <button class="btn btn-primary btn-block" :disabled="loading">
            <span
                v-show="loading"
                class="spinner-border spinner-border-sm"
            ></span>
            <span>Войти</span>
          </button>
        </div>

        <div class="form-group">
          <div v-if="message" class="alert alert-danger" role="alert">
            {{ message }}
          </div>
        </div>
      </Form>
    </div>
  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "Login",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      phoneNumber: yup.string().required("Введите номер телефона!"),
      password: yup.string().required("Введите пароль!"),
    });

    return {
      loading: false,
      message: "",
      schema,
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    },
    created() {
      if (this.loggedIn) {
        this.$router.push("/home");
      }
    }
  },
  methods: {
    handleLogin(user) {
      this.message = "";
      this.loading = true;

      this.$store.dispatch("auth/login", user).then(
          () => {
            this.message = "Success"

            this.$router.push("/home").then(() => {
              window.location.reload()
            })
          },
          (error) => {
            this.loading = false;
            this.message = error.message + ": " + error.response.data.error;
          }
      );
    },
  }
};
</script>

<style scoped>
.card {
  align-items: center;
  align-content: center;
  vertical-align: center;
  background: none;
  border: none white;
}

.card-container {
  background: #f7f7f7;
  padding: 20px;
  border-radius: 100px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  max-width: 800px;
  width: 100%;
  text-align: left;
  align-items: center;
  align-content: center;
  vertical-align: center;
}

.form-group {
  margin-bottom: 15px;
}

.form-label {
  font-weight: bold;
  font-size: 20px;
  text-align: center;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
}

.form-control {
  width: 100%;
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  transition: border-color 0.3s;
}

.form-control:focus {
  border-color: #2575fc;
  outline: none;
  box-shadow: 0 0 5px rgba(37, 117, 252, 0.5);
}

.btn {
  display: inline-block;
  font-size: 1em;
  font-weight: 600;
  text-align: center;
  text-decoration: none;
  padding: 10px 15px;
  margin: 10px 5px;
  border-radius: 5px;
  transition: background-color 0.3s ease;
}

.btn-primary {
  background-color: #007bff;
  color: #fff;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.btn-dark {
  background-color: #343a40;
  color: #fff;
}

.btn-dark:hover {
  background-color: #1d2124;
}

.btn-block {
  width: 100%;
}

.spinner-border {
  margin-right: 5px;
}

.error-feedback {
  color: #ff4d4d;
  font-size: 0.875em;
  margin-top: 5px;
}

.alert {
  margin-top: 20px;
  padding: 10px;
  border-radius: 5px;
  text-align: left;
}
</style>