<template>
  <div class="col-md-12">
    <div class="card card-container">
      <img
          id="profile-img"
          src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
          sizes="(max-width:100px) 20px, 5vw"
          class="profile-img-card"
          alt="Not found"
      />
      <div v-if="!created">
        <Form ref="form" @submit="handleProfile" :validation-schema="schema">
          <div class="form-group">
            <Field type="file" name="image" class="form-control-file" accept="image/jpeg, image/png, image/gif" @change="handleImageUpload" />
            <ErrorMessage name="image" class="text-danger" />
          </div>
          <div class="form-group">
            <label for="documentType">Тип документа, удостоверяющего личность</label>
            <Field name="selectedType" as="select" class="form-control">
              <option value="" disabled selected>Выберите тип документа</option>
              <option v-for="(documentType, index) in documentTypes" :key="index" :value="index">
                {{ documentType }}
              </option>
            </Field>
            <ErrorMessage name="selectedCompany" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="serialNumber">Серийный номер документа</label>
            <Field name="serialNumber" type="text" class="form-control" />
            <ErrorMessage name="serialNumber" class="error-feedback" />
          </div>
          <div class="form-group" v-for="(field, index) in documentFields" :key="index">
            <label :for="'documentField' + index">Поле документа {{ index + 1 }}</label>
            <div class="input-group">
              <input v-model="field.type" type="text" class="form-control" placeholder="Тип поля">
              <input v-model="field.value" type="text" class="form-control" placeholder="Значение поля">
              <div class="input-group-append">
                <button @click="removeField(index)" class="btn btn-outline-secondary" type="button">Удалить</button>
              </div>
            </div>
          </div>
          <button @click="addField" class="btn btn-primary btn-dark mb-1">Добавить поле</button>
          <div class="form-group">
            <button class="btn btn-primary btn-block" :disabled="loading">
              <span v-show="loading" class="spinner-border spinner-border-sm"></span>
              Создать карточку
            </button>
          </div>
        </Form>
      </div>
      <div v-else>
        <div class="card-body">
          <h5 class="card-title">Информация о пользователе</h5>
          <p class="card-text">Номер телефона: {{ profileInfo.phone }}</p>
          <p class="card-text">Должность: {{ profileInfo.post }}</p>
        </div>
      </div>
      <div v-if="message" :class="created ? 'alert-success' : 'alert-danger'" class="alert">
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "Employee",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    return {
      loading: false,
      message: "",
      schema: yup.object().shape({
        image: yup.mixed().required("Загрузите ваше фото или изображение документа, содержащего фото"),
      }),
      documentTypes: [ "Паспорт", "Водительские права" ],
      documentFieldTypes: [ "Дата выдачи", "Выдавший орган" ],
      documentFields: [{ type: "", value: "" }],
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn
    },
    created() {
      return this.$store.state.created;
    },
    profileInfo() {
      return this.$store.state.profile;
    },
  },
  methods: {
    addField() {
      this.documentFields.push({ type: "", value: "" });
    },
    removeField(index) {
      this.documentFields.splice(index, 1);
    },
    handleImageUpload(event) {
      const file = event.target.files[0];
      if (file) {
        this.image = file;
      }
    },
    handleProfile() {
      if (!this.created) {
        this.$refs.form.validate().then(success => {
          if (success) {
            this.fillProfile();
          }
        });
      }
    },
    fillProfile(profile) {
      let user = JSON.parse(localStorage.getItem('user'));

      // FIXME: test data
      const formData = new FormData();
      formData.append('image', this.image);
      formData.append('profileData', new Blob([JSON.stringify({
        serialNumber: "test1",
        documentType: "Паспорт",
        documentFields: [
          {
            "type": "Дата выдачи",
            "value": "OKOK",
          },
          {
            "type": "tetsРПГ",
            "value": "щушс"
          }
        ]
      })], { type: 'application/json' }));

      this.loading = true;

      this.$store.dispatch("employee/fillProfile", formData).then(
          () => {
            this.loading = false;
            this.$store.state.created = true;
            this.$router.push("/profile");
          },
          (error) => {
            this.loading = false;
            if (error.response && error.response.status === 401) {
              this.$store.dispatch('auth/refreshTokens', user).then(
                  response => {
                    this.fillProfile();
                  },
                  (error) => {
                    if (error.response && error.response.status === 401) {
                      this.$store.dispatch('auth/logout');
                      this.$router.push('/login');
                    } else {
                      this.message = error.message + ": " + error.response.data.error;
                    }
                  }
              );
            } else {
              this.message = error.message + ": " + error.response.data.error;
            }
          }
      );
    },
    getProfile() {

    }
  },
};
</script>

<style scoped>
.card-container.card {
  max-width: 700px !important;
  padding: 40px 40px;
  margin: auto;
}

.card {
  background-color: #f7f7f7;
  padding: 20px 25px 30px;
  border-radius: 2px;
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
}

.profile-img-card {
  width: 250px;
  height: 250px;
  margin: auto;
  display: block;
  border-radius: 50%;
}

.error-feedback {
  color: red;
}

.form-group {
  margin-bottom: 20px;
}

.input-group {
  width: 100%;
}

.input-group-append {
  flex: none;
}

.btn-primary {
  width: 100%;
}
</style>
