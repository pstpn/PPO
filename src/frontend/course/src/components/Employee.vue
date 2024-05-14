<template>
  <div class="col-md-12">
    <div class="card card-container">
      <img
          id="profile-img"
          :src="employeePhotoURL"
          sizes="(max-width:100px) 20px, 5vw"
          class="profile-img-card"
          alt="Not found"
      />
      <div v-if="!created">
        <Form ref="form" @submit="handleProfile" :validation-schema="schema">
          <div class="form-group">
            <Field type="file" name="image" class="form-control-file" accept="image/jpeg" @change="handleImageUpload" />
            <ErrorMessage name="image" class="text-danger" />
          </div>
          <div class="form-group">
            <label for="documentType">Тип документа, удостоверяющего личность</label>
            <Field name="selectedType" as="select" class="form-control">
              <option value="" disabled selected>Выберите тип документа</option>
              <option v-for="(documentType, index) in documentTypes" :key="index" :value="documentType">
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
          <div class="form-group">
            <button class="btn btn-primary btn-block" :disabled="loading">
              <span v-show="loading" class="spinner-border spinner-border-sm"></span>
                Создать карточку
            </button>
          </div>
        </Form>
        <button @click="addField" class="btn btn-primary btn-dark mb-1">Добавить поле</button>
      </div>
      <div v-else>
        <div class="card-body">
          <h5 class="card-title">Личная информация</h5>
          <p class="card-text">Тип документа: {{ profileInfo.documentType }}</p>
          <p class="card-text">Серийный номер документа: {{ profileInfo.serialNumber }}</p>
          <table class="table">
            <tbody>
            <tr v-for="(pair, index) in profileInfo.documentFields" :key="index">
              <td>{{ pair.type }}</td>
              <td>{{ pair.value }}</td>
            </tr>
            </tbody>
          </table>
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
    const schema = yup.object().shape({
      image: yup.mixed().required("Загрузите ваше фото или изображение документа, содержащего фото"),
      serialNumber: yup.string().required("Введите серийный номер документа!"),
    });
    return {
      loading: false,
      message: "",
      schema,
      documentTypes: [ "Паспорт", "Водительские права" ],
      documentFieldTypes: [ "Дата выдачи", "Выдавший орган" ],
      documentFields: [],
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn
    },
    created() {
      if (!this.profileInfo) {
        this.$store.dispatch("employee/getProfile").then(
            () => {
              this.$store.dispatch("employee/getEmployeePhoto");
            },
            (error) => {
              if (error.response && error.response.status === 404) {
                this.$store.state.employee.profile = null;
              }
            }
        )
      }
      return this.profileInfo;
    },
    profileInfo() {
      return this.$store.state.employee.profile;
    },
    employeePhotoURL() {
      return this.$store.state.employee.photoURL;
    }
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
    handleProfile(profile) {
      if (!this.created) {
        this.$refs.form.validate().then(success => {
          if (success) {
            this.fillProfile(profile);
          }
        });
      }
    },
    fillProfile(profile) {
      let user = JSON.parse(localStorage.getItem('user'));

      const formData = new FormData();
      formData.append('image', this.image);
      formData.append('profileData', new Blob([JSON.stringify({
        serialNumber: profile.serialNumber,
        documentType: profile.selectedType,
        documentFields: this.documentFields,
      })], { type: 'application/json' }));

      this.loading = true;

      this.$store.dispatch("employee/fillProfile", formData).then(
          () => {
            this.loading = false;
            window.location.reload();
          },
          (error) => {
            this.loading = false;
            if (error.response && error.response.status === 401) {
              this.$store.dispatch('auth/refreshTokens', user).then(
                  response => {
                    this.fillProfile(profile);
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
  width: 200px;
  height: 200px;
  margin: 0 auto 10px;
  display: block;
  -moz-border-radius: 50%;
  -webkit-border-radius: 50%;
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
