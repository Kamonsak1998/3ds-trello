<template>
  <div class="animated fadeIn">
    <div>
      <p>
        <label>
          <b>Select Theme</b>
        </label> &nbsp;
        <select @change="updateTheme">
          <option value="palette1">Theme 1</option>
          <option value="palette2">Theme 2</option>
          <option value="palette3">Theme 3</option>
          <option value="palette4">Theme 4</option>
          <option value="palette5">Theme 5</option>
          <option value="palette6">Theme 6</option>
          <option value="palette7">Theme 7</option>
          <option value="palette8">Theme 8</option>
          <option value="palette9">Theme 9</option>
          <option value="palette10">Theme 10</option>
        </select>
      </p>
      <apexchart type="pie" height="350" :options="chartOptions" :series="series"></apexchart>
    </div>
  </div>
</template>

<script>
export default {
  name: "Pie",
  props: {
      model: {
          required: true
      }
  },
  mounted() {
    this.series = this.model.data
    this.chartOptions = {...this.chartOptions, ...{labels: this.model.name}};
  },
  data: function() {
    return {
        series: [10,20,30,40,50,60,70],
        chartOptions: {
          pie: {
            horizontal: true,
            legend: {
                position: "bottom"
            }
          },
          labels: ['Team A', 'Team B', 'Team C', 'Team D', 'Team E', 'Team E', 'Team E'],
          legend: {
            show: true,
            showForSingleSeries: false,
            showForNullSeries: true,
            showForZeroSeries: true,
            position: 'bottom',
            horizontalAlign: 'center', 
            floating: false,
            fontSize: '12px',
            offsetX: 0,
            offsetY: 0,
        },
        subtitle: {
          text: "Pie Total",
          align: "left",
          margin: 10,
          offsetX: 0,
          offsetY: 0,
          floating: false,
          style: {
              fontSize: "16px",
          }
        },
      }
    };
  },
  methods: {
    updateTheme(e) {
      this.chartOptions = {
        theme: {
          palette: e.target.value
        }
      };
    }
  }
};
</script>
