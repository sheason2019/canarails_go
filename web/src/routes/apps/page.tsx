import { Box, Typography } from '@mui/material';
import NewAppButton from './components/new-app-button';
import AppList from './components/app-list';

export default function Apps() {
  return (
    <main>
      <Typography variant="h4" sx={{ mt: 2 }}>
        Apps
      </Typography>
      <Box sx={{ my: 2 }}>
        <NewAppButton />
      </Box>
      <AppList />
    </main>
  );
}
