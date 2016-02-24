function logout() {
  ret = confirm("ログアウトします、宜しいでしょうか ?");
  if (ret == true){
    location.href = "/";
  }
}