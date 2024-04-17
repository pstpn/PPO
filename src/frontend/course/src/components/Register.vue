<template>
  <div class="col-md-12">
    <div class="card card-container">
      <img
          id="profile-img"
          src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
          class="profile-img-card"
      />
      <Form @submit="handleRegister" :validation-schema="schema">
        <div v-if="!successful">
          <div class="form-group">
            <label for="phoneNumber">Номер телефона</label>
            <Field name="phoneNumber" type="text" class="form-control" />
            <ErrorMessage name="phoneNumber" class="error-feedback" />
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
    // if (this.loggedIn) {
    //   this.$router.push('/home');
    // }
  },
  methods: {
    handleRegister(user) {
      this.message = "";
      this.successful = false;
      this.loading = true;

      console.log(user);

      this.$store.dispatch("auth/register", user).then(
          (data) => {
            this.message = "Success";
            this.successful = true;
            this.loading = false;

            // this.$router.push('/home')
          },
          (error) => {
            this.message =
                (error.response &&
                    error.response.data &&
                    error.response.data.message) ||
                error.message ||
                error.toString();
            this.successful = false;
            this.loading = false;
          }
      );
    },
  },
};
</script>

<style scoped>
label {
  display: block;
  margin-top: 10px;
}

.card-container.card {
  max-width: 350px !important;
  padding: 40px 40px;
}

.card {
  background-color: #f7f7f7;
  padding: 20px 25px 30px;
  margin: 50px auto 25px;
  -moz-border-radius: 2px;
  -webkit-border-radius: 2px;
  border-radius: 2px;
  -moz-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  -webkit-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
}

.profile-img-card {
  width: 96px;
  height: 96px;
  margin: 0 auto 10px;
  display: block;
  -moz-border-radius: 50%;
  -webkit-border-radius: 50%;
  border-radius: 50%;
}

.error-feedback {
  color: red;
}
</style>