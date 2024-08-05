import { Link, Typography } from '@mui/material';
import Dependencies from './components/dependencies';
import VersionData from './components/version-data';

export default function About() {
  return (
    <main>
      <Typography sx={{ mb: 1 }} variant="h5">
        版本信息
      </Typography>
      <VersionData />
      <Dependencies />
      <Typography sx={{ mt: 3, mb: 1 }} variant="h5">
        项目信息
      </Typography>
      <Typography>
        GitHub 仓库地址：
        <Link
          href="https://github.com/sheason2019/canarails_go"
          target="_blank"
          underline="none"
        >
          https://github.com/sheason2019/canarails_go
        </Link>
      </Typography>
    </main>
  );
}
