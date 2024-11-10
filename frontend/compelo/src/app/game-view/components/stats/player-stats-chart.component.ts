import { JsonPipe } from '@angular/common';
import { Component, Input, OnChanges, ViewChild } from '@angular/core';
import { ChartConfiguration, ChartType } from 'chart.js';
import { BaseChartDirective } from 'ng2-charts';

import { PlayerStats } from '@generated/api';

interface DataPoint {
  x: string;
  y: number;
}

@Component({
  selector: 'app-player-stats-chart',
  template: `
    <div class="flex">
      <div class="flex-item">
        <div style="display: block;">
          <canvas
            class="compelo-container"
            baseChart
            [data]="chartData"
            [options]="chartOptions"
            [type]="chartType"
            [legend]="true"
          ></canvas>
        </div>
      </div>
    </div>
  `,
  standalone: true,
  imports: [JsonPipe, BaseChartDirective],
})
export class PlayerStatsChartComponent implements OnChanges {
  @Input()
  players: PlayerStats[];

  @ViewChild(BaseChartDirective, { static: true }) chart: BaseChartDirective;

  chartType: ChartType = 'line';
  chartData: ChartConfiguration<'line', DataPoint[]>['data'] = { datasets: [] };
  chartOptions: ChartConfiguration['options'] = {
    responsive: true,
    aspectRatio: 1.3,
  };

  ngOnChanges() {
    if (this.players) {
      const labels = new Set<string>();
      this.chartData.datasets = this.players.map((player) => {
        return {
          label: player.name,
          data: Object.keys(player.history).map((key) => {
            labels.add(key);
            return {
              x: key,
              y: player.history[key].rating,
            };
          }),
        };
      });
      this.chartData.labels = Array.from(labels);
      this.chart?.update();
    }
  }
}
