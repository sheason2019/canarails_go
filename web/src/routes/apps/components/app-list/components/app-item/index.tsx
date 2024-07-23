import { components } from '@/api/api-gen';
import {
  Button,
  Card,
  CardContent,
  IconButton,
  Stack,
  Typography,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';

interface Props {
  app: components['schemas']['App'];
}

export default function AppItem({ app }: Props) {
  return (
    <Card>
      <CardContent>
        <Typography variant="h5">{app.title}</Typography>
        <Typography variant="body1">简介：{app.description}</Typography>
        <Stack direction="row" justifyContent="space-between" sx={{ mt: 1 }}>
          <IconButton color="error">
            <DeleteIcon />
          </IconButton>
          <Button variant="contained">查看详情</Button>
        </Stack>
      </CardContent>
    </Card>
  );
}
