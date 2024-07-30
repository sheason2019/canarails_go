import Grid from '@mui/material/Unstable_Grid2/Grid2';
import useApp from '../../hooks/use-app';
import AppHostnameCard from './components/app-hostname-card';
import { Alert } from '@mui/material';

export default function AppHostnames() {
  const { data } = useApp();
  const app = data?.data;

  if (!app?.hostnames.length) {
    return (
      <Alert severity="warning">
        请至少配置一个域名匹配，否则 App 将不会应用到 Kubernetes 集群
      </Alert>
    );
  }

  return (
    <Grid container spacing={2}>
      {app?.hostnames.map(hostname => (
        <Grid key={hostname} xs={12} md={4}>
          <AppHostnameCard hostname={hostname} />
        </Grid>
      ))}
    </Grid>
  );
}
