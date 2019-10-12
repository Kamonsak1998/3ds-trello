<template>
  <div class="animated fadeIn">
    <apexchart height="350" type="line" :options="chartOptions" :series="series"></apexchart>
  </div>
</template>

<script>
export default {
  name: "burndownChart",
   mounted: function() {
     this.series[0] = {...this.series[0],...{data: this.model.idealBurn}}
     this.series[1] = {...this.series[1],...{data: this.model.actualBurn}}
     this.chartOptions = {...this.chartOptions, ...{title:{text:this.model.titleSprint}}};
     this.chartOptions = {...this.chartOptions, ...{subtitle:{text:this.model.startDate+ " - " +this.model.endDate}}};
     this.chartOptions = {...this.chartOptions, ...{xaxis: {categories: this.model.datePeriod}}};
     
   },
   props: {
      model: {
          required: true,
      },
  },
  data: function() {
    return {
      series: [
        {
          name: "Ideal Burn",
          color: "rgba(255,0,0,0.25)",
          lineWidth: 2,
          data: []
        },
        {
          name: "Actual Burn",
          color: "rgba(0,120,200,0.75)",
          marker: {
            radius: 6
          },
          data: []
        }
      ],
      chartOptions: {
        plotOptions: {
          line: {
            lineWidth: 3
          },
          tooltip: {
            hideDelay: 200
          }
        },
        title: {
          text: "",
          x: -20,
          align: "center",
          style: {
              fontSize: "20px",
          }
        },
        subtitle: {
            text: "",
            align: "center",
            margin: 10,
            x: -20,
            floating: false,
            style: {
                fontSize: "15px",
            }
        },
        xaxis: {
          categories: []
        },
        yaxis: {
          title: {
            text: "Score"
          },
          plotLines: [
            {
              value: 0,
              width: 1
            }
          ]
        },
        legend: {
          layout: "vertical",
          align: "right",
          verticalAlign: "middle",
          borderWidth: 0
        }
      }
    };
  }
};
</script>
