test:
	echo `pwd`
clearhost:
	ssh-keygen -R "127.0.0.1"
stopsftp: clearhost
	docker stop sftp
stopsftp2: clearhost
	docker stop sftp2
startsftp: clearhost
	docker run --rm --name sftp -v `pwd`/data:/data/incoming -d -p 22:22 writl/sftp
startsftp-pwd: clearhost
	docker run -e USER=test -e PASS=password --name sftp -v `pwd`/data:/data/incoming -d --rm -p 22:22 writl/sftp
startsftp-nopass: clearhost
	docker run -v `pwd`/data:/data/incoming -e USER=test -e PASS=password -e PUBKEY="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDC8troh+aqnGhj1cjLiVEDr+ePYyle0BHZmaQWpgkMG4FPae1BIIuuPDGeG0osWTvPEuTxQbT5ee5WaOlruBmY2RBy3Tg/GQ/MgTKvH55fDEYL5i0NgrUVqu7cVobVQF1UmK//NkMthcZjVuXuUK4yBQP8JXk8XJ7ePnOH9frimSq1R7DGLb3KxRXo1vFOTF5eYH5alOLW+TxgIe0StkPqGBDcK1mIFkCiiSa/nOe4DnNMhcCwjHN9PSwRwlFacT/S7+r4oQtf0114x8X1waQ69pjh7tKtUWgOzTVtz9J/+YP6rWpY+vYSGx8BQclqKCvvRGzdCw1K7tQApkL6ExRr lk@MacBookeLahssen" --name sftp -d --rm -p 22:22 writl/sftp
shellsftp: clearhost
	docker exec -it sftp /bin/bash
shellsftp2: clearhost
	docker exec -it sftp2 /bin/bash
startsftp2:
	docker run --name sftp2 --rm -p 22:22 -d atmoz/sftp test:password:::upload