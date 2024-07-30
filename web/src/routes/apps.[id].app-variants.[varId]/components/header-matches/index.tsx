import Descriptions from '@/common/descriptions';
import { Stack, Box } from '@mui/material';
import DeleteHeaderMatchButton from './delete-header-match-button';
import useAppVariant from '../../hooks/use-app-variant';

export default function HeaderMatches() {
  const { data: appVarData } = useAppVariant();
  const appVar = appVarData?.data;

  if (appVar?.matches.length) {
    return (
      <Descriptions
        items={
          appVar?.matches.map(match => ({
            label: (
              <Stack direction="row" alignItems="center">
                <Box sx={{ flex: 1 }}>{match.header}</Box>
                <DeleteHeaderMatchButton header={match.header} />
              </Stack>
            ),
            value: match.value,
          })) ?? []
        }
      />
    );
  }

  return null;
}
