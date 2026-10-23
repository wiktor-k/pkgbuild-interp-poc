# Maintainer: Jelle van der Waa <jelle@archlinux.org>
# Maintainer: Frederik Schwan <freswa at archlinux dot org>
# Maintainer: Robin Candau <antiz@archlinux.org>
# Contributor: Sander Boom <sanderboom@gmail.com>

pkgname=ansible-lint
pkgver=26.10.1
pkgrel=1
pkgdesc="Checks playbooks for practices and behaviour that could potentially be improved."
arch=('any')
url="https://github.com/ansible/ansible-lint"
license=('GPL-3.0-or-later')
depends=(python ansible-core git yamllint
  python-{ansible-compat,black,cffi,cryptography,distro,filelock,jsonschema,packaging,pathspec,referencing,ruamel-yaml,subprocess-tee,wcmatch,yaml,importlib-metadata})
makedepends=(python-{build,installer,setuptools,setuptools-scm,wheel})
checkdepends=(ansible-creator mypy podman python-jmespath python-pylint python-pytest python-pytest-mock python-uv)
optdepends=('ansible: check official ansible collections')
source=(git+https://github.com/ansible/ansible-lint.git#tag=v$pkgver
        disable_version_check.patch
        ignore_yamllint_unknown_option.patch
        no-galaxy-test-reqs.patch)
b2sums=('7b3879cfb8138559ddbb34224f7ec066796f295be1c395b1139c1cc12e1f69177b3b34dbb843ab0da3e6044e295f132c3f8e7b807e7b359fc8667c019e9e8fde'
        '98294f267ca693c0bc3921f8e076d674a219a891502cd31a0af789bc0b1447b53834b9c85853a134f6bc1ac384f31cb174cba2d55fbcc1636cae9bd3c0bd8f84'
        '8e419c65642bdbd60aa81d1e204139e69fc4c86aaddde4131fbfe3c21f5751e608a5a635d3ec4f518d3d5dc9d254a84b302e09bb4f873e21628e2f014b151516'
        '9074d620b257609e000d1835fcce76b18d6c8b3280b5ab345bf5a98fdcc62b06b01dd3fb146fd2b3853d3568c315b2e5e09f4fe1c4f5586a7ef9ecb989a128e2')

prepare() {
  cd ${pkgname}
  # stop ansible-lint from calling home every 24h to check for a new version
  patch -Np1 < "${srcdir}/disable_version_check.patch"
  # remove yamllint 'forbid-duplicated-merge-keys' unknown rule/option (generates errors during tests)
  patch -Np1 < "${srcdir}/ignore_yamllint_unknown_option.patch"
  # stop trying to pull test reqs from galaxy
  patch -Np1 < "${srcdir}/no-galaxy-test-reqs.patch"
}

build() {
  cd ${pkgname}
  export SETUPTOOLS_SCM_PRETEND_VERSION="${pkgver}"
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd $pkgname
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  local pytest_options=(
    -vv
    --ignore test/test_schemas.py  # requires deprecated spdx
    # don't run tests against venvs and hardcoded locations
    --deselect 'test/test_eco.py::test_eco[bootstrap]'
    --deselect 'test/test_eco.py::test_eco[cisco.nxos]'
    --deselect 'test/test_eco.py::test_eco[colsystem]'
    --deselect 'test/test_eco.py::test_eco[debops]'
    --deselect 'test/test_eco.py::test_eco[docker-rootless]'
    --deselect 'test/test_eco.py::test_eco[hardening]'
    --deselect 'test/test_eco.py::test_eco[mysql]'
    --deselect 'test/test_main.py::test_call_from_outside_venv[normal]'
    --deselect 'test/test_main.py::test_call_from_outside_venv[isolated]'
    --deselect 'test/test_main.py::test_call_from_outside_venv[in]'
    --deselect 'test/test_main.py::test_call_from_outside_venv[missing]'
    --deselect 'test/test_rules_collection.py::test_rules_id_format'
    --deselect 'test/test_main.py::test_get_version_warning[v1.2.2-True-pre-release-1]'
    --deselect 'test/test_main.py::test_get_version_warning[v1.2.3-False--1]'
    --deselect 'test/test_main.py::test_get_version_warning[v1.2.4-True-new release-2]'
    --deselect 'test/test_main.py::test_nodeps[1]'
    --deselect 'test/test_main.py::test_nodeps[2]'
    --deselect 'test/test_main.py::test_broken_ansible_cfg'
    --deselect 'test/test_app.py::test_generate_ignore'
    --deselect 'test/test_app.py::test_app_no_matches'
    --deselect 'test/test_app.py::test_with_inventory_concurrent_syntax_checks'
    --deselect 'test/test_file_utils.py::test_discover_lintables_silent'
    --deselect 'test/test_pattern.py::test_creator_scaffolded_pattern'
    # ignore yamllint tests
    --deselect 'test/test_yaml_utils.py::test_yamllint_incompatible_config'
    --deselect 'test/rules/test_syntax_check.py::test_syntax_check_role'
    --deselect 'test/test_transformer.py::test_transformer[strings]'
    # ignore tests that wants to pull stuff from ansible-galaxy or pip
    --deselect 'test/test_transformer.py::test_pruned_err_after_fix'
    --deselect 'test/test_dependencies_in_meta.py::test_external_dependency_is_ok'
    --deselect 'test/test_file_utils.py::test_kinds[3]'
    --deselect 'test/test_file_utils.py::test_kinds[9]'
    --deselect 'test/test_file_utils.py::test_kinds[10]'
    --deselect 'test/test_file_utils.py::test_kinds[11]'
    --deselect 'test/test_file_utils.py::test_kinds[15]'
    --deselect 'test/test_file_utils.py::test_kinds[23]'
    --deselect 'test/test_include_miss_file_with_role.py::test_cases_warning_message'
    --deselect 'test/test_include_miss_file_with_role.py::test_cases_that_do_not_report[inplace]'
    --deselect 'test/test_include_miss_file_with_role.py::test_cases_that_do_not_report[relative]'
    --deselect 'test/test_config.py::test_guess_install_method_venv'
    --deselect 'test/test_config.py::test_guess_install_method_user_install'
    --deselect 'test/test_config.py::test_guess_install_method_pip_error'
    --deselect 'test/test_config.py::test_guess_install_method_no_uninstall_paths'
    --deselect 'test/test_import_playbook.py::test_task_hook_import_playbook'
    --deselect 'test/test_import_playbook.py::test_import_playbook_from_collection'
    --deselect 'test/test_transformer.py::test_transformer[fqcn]'
   )

  # install to temporary location, as importlib is used
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH"
  export PATH="$PWD/test_dir/usr/bin:$PATH"
  pytest "${pytest_options[@]}" test/
}

package() {
  cd ${pkgname}
  python -m installer --destdir="${pkgdir}" dist/*.whl
}
