import {DataSource} from "typeorm";
import {SystemStatus} from "./entities/system-status.entity";
import _ from "lodash";
import {IncidentStatus} from "./entities/incident-status.entity";

export const dataSource = new DataSource({
  type: "mysql",
  host: "127.0.0.1",
  port: 3308,
  username: "root",
  password: "therootpassword",
  database: "status_for_systems",
  synchronize: false,
  entities: [
    SystemStatus,
    IncidentStatus,
  ],
});

const systemStatusNames = [
  'Operational',
  'Degraded performance',
  'Partial outage',
  'Major outage',
];

const incidentStatusNames = [
  'Investigating',
  'Identified',
  'Monitoring',
  'Resolved',
];

async function addStatuses() {
  // System statuses
  const systemStatusRepository = dataSource.getRepository(SystemStatus);
  const systemStatuses = await systemStatusRepository.find();

  for (const name of systemStatusNames.filter((name) => !_.map(systemStatuses, 'name').includes(name))) {
    const status = new SystemStatus({name});
    await systemStatusRepository.save(status);
  }


  // Incident statuses
  const incidentStatusRepository = dataSource.getRepository(IncidentStatus);
  const incidentStatuses = await incidentStatusRepository.find();

  for (const name of incidentStatusNames.filter((name) => !_.map(incidentStatuses, 'name').includes(name))) {
    const status = new IncidentStatus({name});
    await incidentStatusRepository.save(status);
  }
}

export async function run() {
  await dataSource.initialize();
  console.log('Connected to database');

  await addStatuses();

}

run();

