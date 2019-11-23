import { Component, Input, OnChanges, ViewChild } from '@angular/core';
import { Chart, ChartDataSets, ChartOptions, ChartType } from 'chart.js';
import { PlayerStats } from '@api';
import { BaseChartDirective, Label } from 'ng2-charts';

@Component({
  selector: 'app-player-stats-chart',
  template: `
    <div class="flex">
      <div class="flex-item">
        <div style="display: block;">
          <canvas
            class="compelo-container"
            baseChart
            [datasets]="lineChartData"
            [labels]="lineChartLabels"
            [options]="lineChartOptions"
            [legend]="lineChartLegend"
            [chartType]="lineChartType"
          ></canvas>
        </div>
      </div>
    </div>
  `
})
export class PlayerStatsChartComponent implements OnChanges {
  @Input()
  players: PlayerStats[];

  @ViewChild(BaseChartDirective, { static: true }) chart: BaseChartDirective;

  lineChartData: ChartDataSets[] = [];
  lineChartLabels: Label[] = [];
  lineChartLegend = true;
  lineChartType: ChartType = 'line';
  lineChartOptions: ChartOptions = {
    responsive: true,
    legend: {
      position: 'top'
    },
    aspectRatio: 1.2,
    scales: {
      xAxes: [{ display: false }]
    },
    elements: {
      line: {
        backgroundColor: 'rgba(0, 0, 0, 0)',
        fill: false,
        tension: 0
      }
    }
  };

  ngOnChanges() {
    if (this.players) {
      const labels = new Set<string>();
      const data = this.players.map(player => {
        return {
          label: player.name,
          data: Object.keys(player.history).map(key => {
            labels.add(key);
            return {
              x: key,
              y: player.history[key].rating
            };
          })
        };
      });

      this.lineChartData = data;
      this.lineChartLabels = Array.from(labels);
    }
  }
}
