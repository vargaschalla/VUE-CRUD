<template>
  <div class="container">
    <div class="row">
      <div class="col-sm-10">
        <h1>Person Form</h1>
        <hr />
        <br /><br />
        <alert :message="message" v-if="showMessage"></alert>
        <button
          type="button"
          class="btn btn-success btn-sm"
          v-on:click="back()"
        >
          Cancel
        </button>
        <br /><br />ID= {{ d.ID }}

        <b-form @submit="onSubmit" @reset="onReset" class="w-100">
          <b-form-group
            id="form-title-group"
            label="Nombre:"
            label-for="form-title-input"
          >
            <b-form-input
              id="form-title-input"
              type="text"
              v-model="d.Name"
              required
              placeholder="Enter title"
            >
            </b-form-input>
          </b-form-group>
          <b-form-group
            id="form-author-group"
            label="Edad:"
            label-for="form-author-input"
          >
            <b-form-input
              id="form-author-input"
              type="text"
              v-model="d.Age"
              required
              placeholder="Enter author"
            >
            </b-form-input>
          </b-form-group>

          <b-button-group>
            <b-button type="submit" variant="primary">Submit</b-button>
            <b-button type="reset" variant="danger">Reset</b-button>
          </b-button-group>
        </b-form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Alert from "./Alert.vue";
export default {
  name: "Person",
  data: function () {
    return {
      list: [],
      message: "",
      showMessage: false,
      d: {
        ID: "",
        Name: "",
        Age: "",
      },
    };
  },
  components: {
    alert: Alert,
  },
  methods: {
    makeToast: function (titulo, texto, tipo) {
      this.toastCount++;
      this.$bvToast.toast(texto, {
        title: titulo,
        variant: tipo,
        autoHideDelay: 5000,
        appendToast: true,
      });
    },
    back: function () {
      this.$router.push("/persons");
    },
    create: function (payload) {
      const path = "http://localhost:8080/v1/persons";
      axios
        .post(path, payload)
        .then(() => {
          console.log(payload);
          this.message = "Person added!";
          this.showMessage = true;
          this.makeToast("Hecho", "Paciente creado", "success");
          this.$router.push("/persons");
        })
        .catch((error) => {
          console.log(error);
        });
    },
    initForm: function () {
      //this.d.ID = '';
      this.d.Name = "";
      this.d.Age = "";
    },
    onSubmit: function (evt) {
      evt.preventDefault();
      const payload = {
        name: this.d.Name,
        age: this.d.Age,
      };
      if (this.d.ID > 0) {
        this.update(payload, this.d.ID);
      } else {
        this.create(payload);
      }
      this.initForm();
    },
    onReset: function (evt) {
      evt.preventDefault();
      this.initForm();
    },
    getById: function (id) {
      const path = "http://localhost:8080/v1/persons/" + id;
      console.log(path);
      axios
        .get(path)
        .then((res) => {
          this.d = res.data;
        })
        .catch((error) => {
          // eslint-disable-next-line
          console.error(error);
        });
    },
    update: function (payload, id) {
      const path = `http://localhost:8080/v1/persons/${id}`;
      axios
        .put(path, payload)
        .then(() => {
          console.log(payload);
          this.message = "Person added!";
          this.showMessage = true;
          this.makeToast("Hecho", "Paciente editado", "success");
          this.$router.push("/persons");
        })
        .catch((error) => {
          // eslint-disable-next-line
          console.error(error);
          //this.getBooks();
        });
    },
  },
  created: function () {
    this.d.ID = this.$route.params.id;
    if (this.d.ID > 0) {
      this.getById(this.d.ID);
    }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.hello {
  color: #42b983;
}
</style>