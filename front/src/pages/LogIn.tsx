import { Box, Button, TextField, Typography } from "@mui/material";
import { ChangeEvent, useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";

type LogInInputInfo = {
  name: string;
  pass: string;
}

export default function LogIn() {
  const navigate = useNavigate();
  const title = "ログイン画面"

  const [userName, setUserName] = useState<string>("");
  const handleInputUserName = useCallback((e: ChangeEvent<HTMLInputElement>) => {
    setUserName(e.target.value)
  }, [])

  const [password, setPassword] = useState<string>("");
  const handleInputPassword = useCallback((e: ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value)
  }, [])

  const handleLogInButton = useCallback(() => {
    const inputInfo: LogInInputInfo = {
      name: userName,
      pass: password
    }
    // ログイン情報の確認
    // TODO: 理想は以下の処理としたい
    // 1. APIを叩く
    // 2. ユーザー名でユーザー情報を取得
    // 3. ユーザー名と紐づくpassword等をレスポンスとして取得
    // 4. レスポンスのパスワードと入力されたパスワードを突合
    // 5. 突合結果もOKだったらatomにユーザー情報をセット
    // 6. loginSuccess()を実行
    if ((inputInfo.name === "user" || inputInfo.name === "manager") && inputInfo.pass === "password") {
      console.log("success")
      loginSuccess(inputInfo)
    } else {
      console.log("fail!!!!")
    }
  }, [userName, password])

  // ログイン成功時処理
  const loginSuccess = useCallback((inputInfo: LogInInputInfo) => {
    // userNameによって処理を分岐
    if (inputInfo.name === "manager") {
      console.log("manager login")
    } else {
      navigate("/List")
    }
  }, [])

  return (
    <>
      <Box>
        <Typography>{title}</Typography>
        <TextField
          placeholder="User Name"
          onChange={handleInputUserName}
        />
        <TextField
          placeholder="Password"
          onChange={handleInputPassword}
        />
        <Button
          onClick={handleLogInButton}
        >
          {"LogIn"}
        </Button>
      </Box>
    </>
  )
}
