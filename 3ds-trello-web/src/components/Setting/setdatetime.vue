<template>
  <div class="container">
      <div class="mb-4">
        <date-range-picker v-on:submit="submitted" />
    </div>
  </div>
</template>

<script>
import moment from "moment";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faCalendarAlt, faCaretDown } from "@fortawesome/free-solid-svg-icons";
import BModalDirective from "bootstrap-vue/es/directives/modal/modal";
import DateRangePicker from "@/components/Setting/DateRangePicker";


library.add(faCalendarAlt, faCaretDown);

export default {
  components: { DateRangePicker},
  directives: { "b-modal": BModalDirective },
  data: () => {
    return {
      startDate: moment
        .utc()
        .subtract(1, "month")
        .startOf("month"),
      endDate: moment
        .utc()
        .subtract(1, "month")
        .endOf("month")
        .startOf("day")
    };
  },
  methods: {
    setscore(arr) {
      this.parameter[arr[0]] = arr[1];
    },
    submitted: function() {
      
    },
    submittedModal: function(range) {
      this.startDate = range.startDate;
      this.endDate = "";
      this.closeModal();
    },
    cancelledModal: function() {
      this.closeModal();
    },
    closeModal: function() {
      this.$refs.exampleModal.hide();
    },
    submittedPopover: function(range) {
      this.startDate = range.startDate;
      this.endDate = range.endDate;
      this.closePopover();
    },
    cancelledPopover: function() {
      this.closePopover();
    },
    closePopover: function() {
      this.$refs.examplePopover.$emit("close");
    }
  },
  filters: {
    dateFormat: function(value) {
      return value ? value.format("YYYY-MM-DD") : "";
    }
  }
};
</script>

<style>
.popover {
  max-width: 800px;
}
.col-setting {
  position: relative;
  width: 100%;
  padding-right: 0px;
  padding-left: 0px;
}
</style>
