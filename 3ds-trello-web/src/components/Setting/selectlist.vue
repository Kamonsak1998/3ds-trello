<template>
  <div class="card p-5">
    *First list:
    <b-form-select
      v-model="selectListed[0]"
      @change="select1Change($event)"
      name="select"
      v-validate="'required'"
      :class="{ 'is-invalid': submitted && errors.has('select') }"
    >
      <option v-for="(list,index) in model" :value="list" :key="index">{{list.listName}}</option>
    </b-form-select>
    <div v-if="submitted && errors.has('select')" class="invalid-feedback">{{ errors.first('select') }}</div>
    <br />Second list:
    <b-form-select
      v-model="selectListed[1]"
      @change="select2Change($event)"
      :disabled="select2disable"
    >
      <option v-for="(list,index) in model2" :value="list" :key="index">{{list.listName}}</option>
    </b-form-select>
    <br />Third list:
    <b-form-select v-model="selectListed[2]" :disabled="select3disable">
      <option v-for="(list,index) in model3" :value="list" :key="index">{{list.listName}}</option>
    </b-form-select>
  </div>
</template>

<script>
export default {
  inject:["parentValidator"],
  created() {
    this.$validator = this.parentValidator
  },
  props: {
    model: {
      type: Array,
      required: true
    },
    selectListed: {
      type: Array,
      required: true
    },
    listed: {
      type: Array,
      required: true
    }
  },
  mounted() {
    this.setSelect();
  },
  computed: {
    model2() {
      this.result = this.model.filter(list => {
        return list !== this.selectListed[0];
      });
      return this.result;
    },
    model3() {
      return this.result.filter(list => {
        return list !== this.selectListed[1];
      });
    }
  },
  data() {
    return {
      result: "",
      select2disable: true,
      select3disable: true,
      submitted : false
    };
  },
  methods: {
    formValidate() {
        this.submitted = true;
        return this.$validator.validate()
      },
    setSelect() {
      if (this.listed) {
        this.selectListed[0] = this.listed[0];
        this.selectListed[1] = this.listed[1];
        this.selectListed[2] = this.listed[2];
      }
    },
    select1Change(event) {
      this.select2disable = false;
    },
    select2Change(event) {
      this.select3disable = false;
    }
  }
};
</script>

<style>
</style>
