import useApi from '@/common/use-api';
import useDialog from '@/common/use-dialog';
import { useFormik } from 'formik';
import { useSnackbar } from 'notistack';
import * as Yup from 'yup';
import dayjs from 'dayjs';
import useFieldProperties from '@/routes/hooks/use-field-properties';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  MenuItem,
  Stack,
  TextField,
} from '@mui/material';
import { useEffect } from 'react';
import { components } from '@/api/api-gen';

interface Props {
  onCreate(value: components['schemas']['CreateUserTokenRes']): void;
}

enum ExpiredTimes {
  null = '',
  week = '7天',
  month = '30天',
  quarter = '90天',
  year = '365天',
}

export default function AddUserTokenButton({ onCreate }: Props) {
  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();
  const { isOpen, onClose, onOpen } = useDialog();

  const formik = useFormik({
    initialValues: {
      title: '',
      description: '',
      expiredTime: ExpiredTimes.null,
    },
    validationSchema: Yup.object({
      title: Yup.string().required('令牌名称不能为空'),
      description: Yup.string(),
      expiredTime: Yup.string()
        .required('有效时间不能为空')
        .test({
          name: 'is-expired-time',
          test(value, ctx) {
            switch (value) {
              case ExpiredTimes.week:
              case ExpiredTimes.month:
              case ExpiredTimes.quarter:
              case ExpiredTimes.year:
                return true;
              default:
                return ctx.createError({ message: 'invalid expired time' });
            }
          },
        }),
    }),
    onSubmit: async values => {
      let expiredAt = dayjs();

      switch (values.expiredTime) {
        case ExpiredTimes.week:
          expiredAt = expiredAt.add(7, 'day');
          break;
        case ExpiredTimes.month:
          expiredAt = expiredAt.add(30, 'day');
          break;
        case ExpiredTimes.quarter:
          expiredAt = expiredAt.add(90, 'day');
          break;
        case ExpiredTimes.year:
          expiredAt = expiredAt.add(365, 'day');
          break;
        default:
          throw new Error('Unexpected ExpiredTimes');
      }

      const res = await api.POST('/api/user-token', {
        body: {
          id: 0,
          title: values.title,
          description: values.description,
          expiredAt: expiredAt.unix() * 1000,
          lastUsedAt: 0,
        },
      });
      const data = res.data;
      if (data) {
        onCreate(data);
      }

      enqueueSnackbar('新增权限令牌成功', { variant: 'success' });
      onClose();
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  useEffect(() => {
    formik.resetForm();
  }, [isOpen]);

  return (
    <>
      <Button variant="contained" onClick={onOpen}>
        新增
      </Button>
      <Dialog open={isOpen} onClose={onClose} fullWidth>
        <form onSubmit={formik.handleSubmit}>
          <DialogTitle>新增权限令牌</DialogTitle>
          <DialogContent>
            <Stack spacing={2}>
              <TextField
                label="令牌名称"
                variant="standard"
                {...getFieldProperties('title')}
              />
              <TextField
                label="有效时间"
                variant="standard"
                select
                {...getFieldProperties('expiredTime')}
              >
                {[
                  ExpiredTimes.week,
                  ExpiredTimes.month,
                  ExpiredTimes.quarter,
                  ExpiredTimes.year,
                ].map(item => (
                  <MenuItem key={item} value={item}>
                    {item}
                  </MenuItem>
                ))}
              </TextField>
              <TextField
                label="令牌简介"
                variant="standard"
                multiline
                minRows={3}
                maxRows={5}
                {...getFieldProperties('description')}
              />
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button onClick={onClose}>取消</Button>
            <Button variant="contained" type="submit">
              提交
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </>
  );
}
