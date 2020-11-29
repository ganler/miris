import subprocess

def f(ds_list, x_list):
	for ds in ds_list:
		for x in x_list:
			subprocess.call(['mkdir', '-p', 'data/{}/frames/{}/'.format(ds, x)])
			subprocess.call(['ffmpeg', '-i', 'data/{}/videos/{}.mp4'.format(ds, x), '-s', '960x540', '-q:v', '1', 'data/{}/frames/{}/%06d.jpg'.format(ds, x)])

# f(['shibuya'], range(6))
# f(['shibuya', 'beach', 'warsaw', 'jacksontown'], range(6))
f(['uav'], ['0006', '0007', '0008', '0009', '0011'])
