import Grid from '@mui/material/Unstable_Grid2/Grid2';
import useApp from '../../hooks/use-app';
import AppHostnameCard from './components/app-hostname-card';
import { Box, Typography } from '@mui/material';

export default function AppHostnames() {
  const { data } = useApp();
  const app = data?.data;

  if (!app?.hostnames.length) {
    return (
      <Box
        sx={{
          height: 48,
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <Typography variant="body2" color="gray">
          暂无数据
        </Typography>
      </Box>
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
