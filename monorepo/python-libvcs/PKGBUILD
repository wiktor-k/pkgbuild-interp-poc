# Maintainer: Andreas 'Segaja' Schleifer <segaja at archlinux dot org>

_name='libvcs'
pkgname="python-${_name}"
pkgver=0.46.0
pkgrel=1
pkgdesc='Lite, typed, pythonic utilities for git, svn, mercurial, etc'
arch=('any')
url='https://libvcs.git-pull.com/'
license=('MIT')
depends=(
  python
)
makedepends=(
  python-build
  python-installer
)
checkdepends=(
  git
  python-gp-libs
  mercurial
  python-hatchling
  python-pytest
  python-pytest-mock
  python-pytest-rerunfailures
  subversion
)
optdepends=(
  'git: for git repository support'
  'mercurial: for mercurial repository support'
  'subversion: for subversion repository support'
  'python-pytest: for pytest plugin'
)
source=("https://github.com/vcs-python/libvcs/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('362beeb84d32a6775d8ec3ff3845cd5915eaf9725e38d80ca31b71422065abc78e684642abc76d9319d3fcba74c9c05fe844da9ee96ce4855d217dee50a5d2d9')
b2sums=('0648ba9ebb672d6631c9238d0e671689995ac6e3b2b6a22c3d76e788f789ae6c92ebcfd0bf8ea5d1d79037b57fa1e20e38201a4496768b2787f75f19af2eaba8')

build() {
  cd "${_name}-${pkgver}"

  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd "${_name}-${pkgver}"

  export GIT_CONFIG_GLOBAL="$HOME/.gitconfig"
  git config --global user.email "custom_email@testemail.com"
  git config --global user.name "Your Name"

  local _site_packages="$(python -c "import site; print(site.getsitepackages()[0])")"

  python -m installer --destdir=test_dir dist/*.whl

  export PYTHONPATH="test_dir${_site_packages}:${PYTHONPATH}"

  pytest -vv tests
}

package() {
  cd "${_name}-${pkgver}"

  python -m installer --destdir="${pkgdir}" dist/*.whl

  install --verbose -D --mode=0644 LICENSE --target-directory "${pkgdir}/usr/share/licenses/${pkgname}"
  install --verbose -D --mode=0644 *.md CHANGES MIGRATION --target-directory "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: tabstop=2 shiftwidth=2 expandtab:
