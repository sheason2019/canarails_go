import { api } from '@/api/api';
import useToken from '@/common/user/use-token';
import useFieldProperties from '@/routes/hooks/use-field-properties';
import { LoadingButton } from '@mui/lab';
import { CardContent, Stack, TextField } from '@mui/material';
import { useFormik } from 'formik';
import { useSnackbar } from 'notistack';
import * as Yup from 'yup';

export default function LoginCard() {
  const { setToken } = useToken();
  const { enqueueSnackbar } = useSnackbar();

  const formik = useFormik({
    initialValues: {
      username: '',
      password: '',
    },
    validationSchema: Yup.object({
      username: Yup.string().required('必填项'),
      password: Yup.string().required('必填项'),
    }),
    onSubmit: async values => {
      const res = await api.POST('/api/auth/login', {
        body: values,
      });
      const token = res.data?.token;
      if (token) {
        setToken(token);
        enqueueSnackbar('登录成功', { variant: 'success' });
      }
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  return (
    <CardContent component="form" onSubmit={formik.handleSubmit}>
      <Stack gap={1} sx={{ px: 1 }}>
        <p className="text-lg text-center font-bold">用户登录</p>
        <TextField
          fullWidth
          label="用户名"
          variant="standard"
          {...getFieldProperties('username')}
        />
        <TextField
          fullWidth
          type="password"
          label="密码"
          variant="standard"
          {...getFieldProperties('password')}
        />
        <div className="mt-3">
          <LoadingButton
            className="w-full"
            type="submit"
            variant="contained"
            loading={formik.isSubmitting}
            loadingIndicator="正在登录"
          >
            登录
          </LoadingButton>
        </div>
      </Stack>
    </CardContent>
  );
}
