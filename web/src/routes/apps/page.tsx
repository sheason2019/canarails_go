import { Box, Typography } from '@mui/material';
import NewAppButton from './components/new-app-button';

export default function Apps() {
  return (
    <main>
      <Typography variant="h4" sx={{ mt: 2 }}>
        Apps
      </Typography>
      <Box sx={{ mt: 2 }}>
        <NewAppButton />
      </Box>
    </main>
  );
}
