<template>
  <div class="card">
    <div class="card card-container">
      <Form @submit="handleRegister" :validation-schema="schema">
        <div v-if="!successful">
          <div class="form-label">
            Регистрация
          </div>
          <div class="form-group">
            <label for="name">Имя</label>
            <Field name="name" type="text" class="form-control" />
            <ErrorMessage name="name" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="surname">Фамилия</label>
            <Field name="surname" type="text" class="form-control" />
            <ErrorMessage name="surname" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="phoneNumber">Номер телефона</label>
            <Field name="phoneNumber" type="text" class="form-control" />
            <ErrorMessage name="phoneNumber" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="company">Компания</label>
            <Field name="selectedCompany" as="select" class="form-control">
              <option value="" disabled selected>Выберите компанию</option>
              <option v-for="(company, index) in companies" :key="index" :value="index">
                {{ company }}
              </option>
            </Field>
            <ErrorMessage name="selectedCompany" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="post">Должность</label>
            <Field name="post" type="text" class="form-control" />
            <ErrorMessage name="post" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="dateOfBirth">Дата рождения (в формате "дд.мм.гггг")</label>
            <Field name="dateOfBirth" type="text" class="form-control" />
            <ErrorMessage name="dateOfBirth" class="error-feedback" />
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
              Зарегистрироваться
            </button>
          </div>
        </div>
      </Form>

      <div
          v-if="message"
          class="alert"
          :class="successful ? 'alert-success' : 'alert-danger'"
      >
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

const phoneRegExp = /^([+]?[0-9\s-()]{3,25})*$/i

export default {
  name: "Register",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      phoneNumber: yup
          .string()
          .required("Введите номер телефона")
          .matches(phoneRegExp, "Некорректный номер телефона")
          .min(11, "Некорректный номер телефона"),
      name: yup
          .string()
          .required("Введите имя"),
      surname: yup
          .string()
          .required("Введите фамилию"),
      post: yup
          .string()
          .required("Введите должность"),
      dateOfBirth: yup
          .string()
          .required("Введите дату рождения")
          .matches(/^\d{2}\.\d{2}\.\d{4}$/, "Некорректный формат даты (дд.мм.гггг)"),
      password: yup
          .string()
          .required("Введите пароль")
    });

    return {
      successful: false,
      loading: false,
      message: "",
      companies: [ "Yandex", "Полисофт", "Ситисофт" ],
      schema,
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    },
  },
  mounted() {
    if (this.loggedIn) {
      this.$router.push('/home');
    }
  },
  methods: {
    handleRegister(user) {
      this.message = "";
      this.successful = false;
      this.loading = true;

      this.$store.dispatch("auth/register", user).then(
          () => {
            this.message = "Success";
            this.successful = true;
            this.loading = false;

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
  },
};
</script>

<style scoped>
.card {
  align-items: center;
  align-content: center;
  vertical-align: center;
  border-color: white;
}

.card-container {
  background: #f7f7f7;
  padding: 20px;
  border-radius: 100px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  max-width: 1000px;
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

.btn:hover {
  background: #1a5bb8;
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