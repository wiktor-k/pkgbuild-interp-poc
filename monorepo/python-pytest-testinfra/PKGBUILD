# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: David Runge <dvzrv@archlinux.org>
# Contributor: Christian Rebischke <chris.rebischke@archlinux.org>
# Contributor: Lex Black <autumn-wind@web.de>
# Contributor: gardar <aur@gardar.net>
# Contributor: Morten Linderud <morten@linderud.pw>

pkgname=python-pytest-testinfra
pkgver=10.2.2
pkgrel=2
pkgdesc='Testinfra test your infrastructures'
arch=(any)
url='https://github.com/pytest-dev/pytest-testinfra'
license=(Apache-2.0)
depends=(
  python
  python-pytest
)
makedepends=(
  git
  python-build
  python-installer
  python-hatchling
  python-hatch-vcs
  python-sphinx
  python-wheel
)
checkdepends=(
  ansible
  python-paramiko
  python-pytest-xdist
  python-pywinrm
  python-tornado
)
optdepends=(
  'ansible: for tests using ansible inventories'
  'docker: for test running docker containers'
  'kubectl: for test running containers in kubernetes'
  'lxd: for test running LXC or LXD containers'
  'podman: for test running podman containers'
  'python-paramiko: for testing remote hosts'
  'python-pywinrm: for testing on Windows hosts'
  'salt: for tests using saltstack'
)
source=("$pkgname::git+$url#tag=$pkgver")
sha512sums=('5b61f4b8322aed2303cab4aa327b5f9f5c63786de5f3a31b00ab9203cda477b8245d4f81a9e5b8555a4092e2aa357c2c854e6be7fff61e99e8ab47fa0bbe6244')
b2sums=('0f4d8fb9164d02315f0cf3477d5d0c9d72e7475292a05b0aab4bfdef89034c1c948f9e158d63f6693bec301d5a8e528eba10794ccd0acdeb442bef817a027a70')

build() {
  cd "$pkgname"

  export SETUPTOOLS_SCM_PRETEND_VERSION="$pkgver"

  # build wheel
  python -m build --wheel --no-isolation

  # install to temporary directory
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  python -m installer --destdir="$PWD/tmp_install" dist/*.whl

  # build man page
  PYTHONPATH="$PWD/tmp_install$site_packages" make -C doc man
}

check() {
  local pytest_options=(
    -vv
  )

  cd "$pkgname"

  # install to temporary location
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  python -m installer --destdir=test_dir dist/*.whl

  # We don't have salt in the repos
  pytest_options=(
    -k 'not test_backend_importables'
  )

  PYTHONPATH="test_dir/$site_packages:$PYTHONPATH" pytest "${pytest_options[@]}"
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # man pages
  install -vDm644 -t "$pkgdir/usr/share/man/man1" doc/build/man/*.1

  # documentation
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" {CHANGELOG,CONTRIBUTING,README}.rst
}
