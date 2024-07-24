import Grid from '@mui/material/Unstable_Grid2/Grid2';
import useApp from '../../hooks/use-app';
import { Card, CardContent } from '@mui/material';

export default function AppHostnames() {
  const { data } = useApp();
  const app = data?.data;

  return (
    <Grid container spacing={2}>
      {app?.hostnames.map(hostname => (
        <Grid key={hostname} xs={12} md={4}>
          <Card>
            <CardContent>{hostname}</CardContent>
          </Card>
        </Grid>
      ))}
    </Grid>
  );
}
