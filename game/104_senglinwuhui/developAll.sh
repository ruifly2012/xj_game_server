

echo "=============开始上传=============="

for (( i = 1; i < 5; i++ )); do
  pwd
  ./develop.sh $i
  echo "============== $i 完成上传=================="
done

echo "==========104_senglinwuhui 上传完毕============"